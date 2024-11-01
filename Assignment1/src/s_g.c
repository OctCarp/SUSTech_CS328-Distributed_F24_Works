#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAT_SIZE (500)
#define EPS (1e-6)

#define abs(x) ((x) > 0 ? (x) : (-(x)))

void init_rand_mat_ptr(double *mat) {
    /* initialize a matrix with random values */
    for (int i = 0; i < MAT_SIZE; ++i) {
        for (int j = 0; j < MAT_SIZE; ++j) {
            mat[i * MAT_SIZE + j] = rand() + (double) rand() / RAND_MAX;
        }
    }
}

void brute_force_matmul_ptr(double *mat1, double *mat2,
                            double *res) {
    /* matrix multiplication of mat1 and mat2, store the result in res */
    for (int i = 0; i < MAT_SIZE; ++i) {
        for (int j = 0; j < MAT_SIZE; ++j) {
            res[i * MAT_SIZE + j] = 0;
            for (int k = 0; k < MAT_SIZE; ++k) {
                res[i * MAT_SIZE + j] += mat1[i * MAT_SIZE + k] * mat2[k * MAT_SIZE + j];
            }
        }
    }
}

int checkRes_ptr(const double *target, const double *res) {
    /* check whether the obtained result is the same as the intended target; if true return 1, else return 0 */
    for (int i = 0; i < MAT_SIZE; ++i) {
        for (int j = 0; j < MAT_SIZE; ++j) {
            double diff = target[i * MAT_SIZE + j] - res[i * MAT_SIZE + j];
            if (abs(diff) > EPS) {
                return 0;
            }
        }
    }
    return 1;
}

int main(int argc, char *argv[]) {
    int rank;
    int mpiSize;
    srand(time(NULL));

    // double a[MAT_SIZE][MAT_SIZE],    /* matrix A to be multiplied */
    // b[MAT_SIZE][MAT_SIZE],       /* matrix B to be multiplied */
    // c[MAT_SIZE][MAT_SIZE],       /* result matrix C */
    // bfRes[MAT_SIZE][MAT_SIZE];   /* brute force result bfRes */

    double *a = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *b = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *c = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *bfRes = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));

    /* You need to intialize MPI here */
    MPI_Init(NULL, NULL);

    // Get the number of processes
    MPI_Comm_size(MPI_COMM_WORLD, &mpiSize);

    // Get the rank of the process
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    int local_size = MAT_SIZE / mpiSize;
    double *local_a = (double *) malloc(local_size * MAT_SIZE * sizeof(double));
    double *local_c = (double *) malloc(local_size * MAT_SIZE * sizeof(double));

    if (rank == 0) {
        /* root */

        /* First, fill some numbers into the matrix */
        // for (int i = 0; i < MAT_SIZE; ++i)
        //     for (int j = 0; j < MAT_SIZE; ++j) {
        //         a[i * MAT_SIZE + j] = i + j;
        //         b[i * MAT_SIZE + j] = i * j;
        //     }

        init_rand_mat_ptr(a);
        init_rand_mat_ptr(b);

        /* Measure start time */
        double start = MPI_Wtime();

        /* Send matrix data to the worker tasks */
        MPI_Scatter(a, local_size * MAT_SIZE, MPI_DOUBLE, local_a, local_size * MAT_SIZE, MPI_DOUBLE, 0,
                    MPI_COMM_WORLD);
        MPI_Bcast(b, MAT_SIZE * MAT_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

        /* Compute its own piece */
        for (int i = 0; i < local_size; ++i) {
            int pos = i * MAT_SIZE;
            for (int j = 0; j < MAT_SIZE; ++j) {
                double temp = 0.0;
                for (int k = 0; k < MAT_SIZE; ++k) {
                    temp += local_a[pos + k] * b[k * MAT_SIZE + j];
                }
                local_c[pos + j] = temp;
            }
        }

        /* Receive results from worker tasks */

        MPI_Gather(local_c, local_size * MAT_SIZE, MPI_DOUBLE, c, local_size * MAT_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

        int local_size_rem = MAT_SIZE % mpiSize;
        for (int i = 0; i < local_size_rem; ++i) {
            int pos = (mpiSize * local_size + i) * MAT_SIZE;
            for (int j = 0; j < MAT_SIZE; ++j) {
                double temp = 0.0;
                for (int k = 0; k < MAT_SIZE; ++k) {
                    temp += a[pos + k] * b[k * MAT_SIZE + j];
                }
                c[pos + j] = temp;
            }
        }

        /* Measure finish time */
        double finish = MPI_Wtime();

        /* Compare results with those from brute force */
        brute_force_matmul_ptr(a, b, bfRes);
        if (!checkRes_ptr(bfRes, c)) {
            printf("ERROR: Your calculation is not the same as the brute force result, please check!\n");
        } else {
            printf("%.10lf\n", finish - start);
        }
    } else {
        /* worker */
        /* Receive data from root and compute, then send back to root */
        MPI_Scatter(a, local_size * MAT_SIZE, MPI_DOUBLE, local_a, local_size * MAT_SIZE, MPI_DOUBLE, 0,
                    MPI_COMM_WORLD);
        MPI_Bcast(b, MAT_SIZE * MAT_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);

        for (int i = 0; i < local_size; ++i) {
            int pos = i * MAT_SIZE;
            for (int j = 0; j < MAT_SIZE; ++j) {
                double temp = 0.0;
                for (int k = 0; k < MAT_SIZE; ++k) {
                    temp += local_a[pos + k] * b[k * MAT_SIZE + j];
                }
                local_c[pos + j] = temp;
            }
        }

        MPI_Gather(local_c, local_size * MAT_SIZE, MPI_DOUBLE, c, local_size * MAT_SIZE, MPI_DOUBLE, 0, MPI_COMM_WORLD);
    }
    free(a);
    free(b);
    free(c);
    free(bfRes);
    free(local_a);
    free(local_c);

    /* Don't forget to finalize your MPI application */
    MPI_Finalize();

    return 0;
}
