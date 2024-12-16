import pandas as pd
import matplotlib.pyplot as plt
import os

CHOSEN_DATE: str = '2018-09-01'
plt.rcParams['font.sans-serif'] = ['SimHei']

def get_path(filename: str) -> str:
    return os.path.join(os.path.dirname(os.getcwd()), 'data', filename)


def read_csv() -> pd.DataFrame:
    r5 = pd.read_csv(get_path("r5.csv"))
    r5['start_time'] = pd.to_datetime(r5['start_time'])
    r5['end_time'] = pd.to_datetime(r5['end_time'])

    return r5


def choose_section(r5: pd.DataFrame) -> tuple[str, str, str]:
    sections = r5.groupby('section')['percentage'].mean().sort_values(ascending=False)
    selected_sections = (sections.index[0], sections.index[len(sections) // 3], sections.index[len(sections) // 2])
    return selected_sections


def plot(r5: pd.DataFrame, selected_sections: tuple[str, str, str]) -> None:
    plt.figure(figsize=(15, 8))

    colors = ('red', 'green', 'blue')
    for i, section in enumerate(selected_sections):
        section_data: pd.Series = r5[r5['section'] == section]
        daily_data: pd.Series = section_data[section_data['start_time'].dt.strftime('%Y-%m-%d') == CHOSEN_DATE]
        final_data: pd.Series = daily_data[['start_time', 'end_time', 'percentage']]

        plt.plot(final_data['start_time'],
                 final_data['percentage'],
                 marker='o',
                 color=colors[i],
                 label=section)

    plt.title(f'Three Sections Percentage in {CHOSEN_DATE}')
    plt.xlabel('Time')
    plt.ylabel('Percentage(%)')
    plt.grid(True)
    plt.legend()
    plt.xticks(rotation=45)

    plt.tight_layout()

    plt.savefig(get_path('parking_utilization.png'))
    plt.show()


if __name__ == '__main__':
    result5 = read_csv()
    selected = choose_section(result5)
    plot(result5, selected)
