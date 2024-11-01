import argparse
import csv
from collections import defaultdict
import matplotlib.pyplot as plt

CSV_PATH = './test_record.csv'
OUTPUT_PATH = './bin/visual_latest.png'


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Plot the average running time for each filename by processes')
    parser.add_argument("--comp", type=str, default="0",
                        help='Compare the local and distributed env, 1 for yes 0 for not')    # 设置默认值为字符 0，不设置默认值则为 None

    args = parser.parse_args()
    
    data = []
    
    with open(CSV_PATH, newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            data.append({
                'Filename': row['Filename'],
                'Proc': int(row['Proc']),
                'Running Time': float(row['Running Time']),
                'Time': row['Time'],
                'Iterations': int(row['Iterations'])
            })

    filtered_data = defaultdict(list)

    if args.comp == '1':
        print('draw the comparison between local and distributed env')
        for row in data:
            if (row['Iterations'] >= 50 and row['Filename']=='sv_gv.c') or (row['Filename'][:5]=='distr'):
                filtered_data[(row['Filename'], row['Proc'])].append(row['Running Time'])
    elif args.comp == '0':
        print('draw the average running time for each local filename by processes')
        for row in data:
            if row['Iterations'] >= 50:
                filtered_data[(row['Filename'], row['Proc'])].append(row['Running Time'])

    average_data = defaultdict(dict)

    for (filename, processes), running_times in filtered_data.items():
        avg_running_time = sum(running_times) / len(running_times)
        average_data[filename][processes] = avg_running_time

    plt.figure(figsize=(10, 6))
    
    for filename, procs_data in average_data.items():
        procs = sorted(procs_data.keys())
        plt.xticks(procs)
        running_times = [procs_data[proc] for proc in procs]
        plt.plot(procs, running_times, marker='o', label=filename)
    
    plt.xlim(left=0)
    plt.ylim(bottom=0)

    plt.xlabel('Processes')
    plt.ylabel('Average Running Time')
    plt.title('Average Running Time for Each Filename by Processes')
    plt.legend(title='Filename')
    
    plt.tight_layout()
    plt.savefig(OUTPUT_PATH)
    plt.show()
    
