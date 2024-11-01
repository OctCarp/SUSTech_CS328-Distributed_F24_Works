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
    /* 
      check whether the obtained result is the same as the intended target; 
      if true return 1, else return 0 
    */
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

    /* You need to intialize MPI here */
    MPI_Init(NULL, NULL);

    // Get the number of processes
    MPI_Comm_size(MPI_COMM_WORLD, &mpiSize);

    // Get the rank of the process
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    // double a[MAT_SIZE][MAT_SIZE],    /* matrix A to be multiplied */
    // b[MAT_SIZE][MAT_SIZE],       /* matrix B to be multiplied */
    // c[MAT_SIZE][MAT_SIZE],       /* result matrix C */
    // bfRes[MAT_SIZE][MAT_SIZE];   /* brute force result bfRes */

    double *a = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *b = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *c = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));
    double *bfRes = (double *) malloc(MAT_SIZE * MAT_SIZE * sizeof(double));

    int quo_sz = MAT_SIZE / mpiSize;
    int rem_cnt = MAT_SIZE % mpiSize;

    double *local_a = (double *) malloc((quo_sz + 1) * MAT_SIZE * sizeof(double));
    double *local_c = (double *) malloc((quo_sz + 1) * MAT_SIZE * sizeof(double));

    int local_szs[mpiSize], displs[mpiSize], scounts[mpiSize];

    for (int i = 0; i < rem_cnt; ++i) {
        local_szs[i] = quo_sz + 1;
    }
    for (int i = rem_cnt; i < mpiSize; ++i) {
        local_szs[i] = quo_sz;
    }

    int offset = 0;
    for (int i = 0; i < mpiSize; ++i) {
        displs[i] = offset;
        offset += local_szs[i] * MAT_SIZE;
        scounts[i] = local_szs[i] * MAT_SIZE;
    }

    double start, finish;
    int root = 0;
    if (rank == root) {
        /* root */

        /* First, fill some numbers into the matrix */

        // for (int i = 0; i < MAT_SIZE; ++i)
        //     for (int j = 0; j < MAT_SIZE; ++j) {
        //         a[i * MAT_SIZE + j] = i + j;
        //         b[i * MAT_SIZE + j] = i * j;
        //     }

        init_rand_mat_ptr(a);
        init_rand_mat_ptr(b);
    }

    //MPI_Barrier(MPI_COMM_WORLD);
    if (rank == root) {
        /* Measure start time */
        start = MPI_Wtime();
    }

    /* Send matrix data to the worker tasks */
    MPI_Scatterv(a, scounts, displs, MPI_DOUBLE, local_a, scounts[rank], MPI_DOUBLE, root, MPI_COMM_WORLD);
    MPI_Bcast(b, MAT_SIZE * MAT_SIZE, MPI_DOUBLE, root, MPI_COMM_WORLD);

    /* Compute its own piece */
    int local_size = local_szs[rank];
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
    MPI_Gatherv(local_c, scounts[rank], MPI_DOUBLE, c, scounts, displs, MPI_DOUBLE, root, MPI_COMM_WORLD);

    //MPI_Barrier(MPI_COMM_WORLD);
    if (rank == root) {
        /* Measure finish time */
        finish = MPI_Wtime();

        /* Compare results with those from brute force */
        brute_force_matmul_ptr(a, b, bfRes);
        if (!checkRes_ptr(bfRes, c)) {
            printf("ERROR: Your calculation is not the same as the brute force result, please check!\n");
        } else {
            printf("%.10lf\n", finish - start);
        }
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
