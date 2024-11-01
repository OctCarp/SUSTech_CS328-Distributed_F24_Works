#!/bin/bash

# Get all .c files in the current directory
c_files=(*.c)

record_path="test_record.csv"
if [ ! -f "$record_path" ]; then
    # if not, create it
    touch "$record_path"
    echo "File $record_path not found, created."
    echo "Filename,Proc,Running Time,Time,Iterations" > "$record_path"
fi

# Check if the output directory exists
mpi_bin_directory="mpi_files"
if [ ! -d "$mpi_bin_directory" ]; then
    # if not, create it
    mkdir -p "$mpi_bin_directory"
    echo "Directory $mpi_bin_directory not found, created."
fi

# Check if there are any .c files
if [[ ${#c_files[@]} -eq 0 ]]; then
    echo "No .c files found in the current directory."
    exit 1
fi

# Display .c file list and prompt the user to choose one
echo "Please select a .c file to compile:"
for i in "${!c_files[@]}"; do
    echo "$i) ${c_files[$i]}"
done

# Read the user's choice
read -p "Enter the file number: " file_index

# Validate the user's input
if [[ ! $file_index =~ ^[0-9]+$ ]] || [[ $file_index -ge ${#c_files[@]} ]]; then
    echo "Invalid input. Please enter a valid file number."
    exit 1
fi

# Get the filename based on the user's choice
filename=${c_files[$file_index]}
bin_filename=$mpi_bin_directory/${filename%.*}.out

# Compile the selected .c file
mpicc "$filename" -o "$bin_filename"
if [[ $? -ne 0 ]]; then
    echo "Compilation failed!"
    exit 1
fi

# Prompt for the number of cores and iterations
# read -p "Enter the number of cores to allocate: " core_num
read -p "Enter the number of iterations: " loop_num

# Run the program loop_num times

for core_num in 1 2 4 8 16 32; do
    # Initialize variables
    total=0.0
    count=0
    echo "Running complied '$filename' for $loop_num times with $core_num cores, waiting..."
        
    for ((i=1; i<=loop_num; i++)); do

        # Run loop_test using mpirun, limiting execution time to 1 second
        result=$(timeout 5s mpirun -np "$core_num" -oversubscribe "./$bin_filename")

        # Check if the program executed successfully and returned a valid float
        if [[ $? -eq 0 ]] && [[ $result =~ ^[0-9]+(\.[0-9]+)?$ ]]; then
            total=$(echo "$total + $result" | bc)
            count=$((count + 1))
        else
            echo "Iteration $i failed or produced invalid output"
            echo "Outpute is: '$result'"
            exit 1
        fi

        # Every 10 iterations, print a status update
        if (( i % 10 == 0 )); then
            echo "Completed $i iterations..."
            average=$(echo "$total / $i" | bc -l)
            echo "Average value: $average"
        fi
    done

    # Calculate and display the average result
    if [[ $count -ne 0 ]]; then
        average=$(echo "$total / $count" | bc -l)
        echo "Finish complied '$bin_filename' for $loop_num times with $core_num cores,"
        echo "Final average value: $average"

        # Get current time "YY-MM-DD HH:MM:SS"
        current_time=$(date +"%y-%m-%d %H:%M:%S")
        # Write the record to the CSV file
        echo "$filename,$core_num,$average,$current_time,$loop_num" >> "$record_path"
        if [[ $? -ne 0 ]]; then
            echo "Failed to save the record to $record_path"
            exit 1
        fi
        echo "Saved the record to $record_path"
    else
        echo "No valid results to calculate an average."
    fi
done
