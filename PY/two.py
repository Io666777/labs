import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.datasets import fetch_california_housing

# ============================================
# ЗАГРУЗКА ДАННЫХ (ВАРИАНТ 12)
# ============================================
data = fetch_california_housing(as_frame=True)
df = data.frame
X = df['AveBedrms'].values

# Параметры варианта 12
trim_percent = 10
wins_percent = 20
B = 2500

print("=" * 70)
print("ВАРИАНТ 12: X = AveBedrms (среднее число спален на дом)")
print("=" * 70)

# ============================================
# ЗАДАНИЕ 1. ДИАГНОСТИКА ПЕРЕМЕННОЙ X
# ============================================
print("\n" + "=" * 70)
print("ЗАДАНИЕ 1. Диагностика переменной X")
print("=" * 70)

mean_x = np.mean(X)
std_x = np.std(X, ddof=1)
median_x = np.median(X)
q1_x = np.quantile(X, 0.25)
q3_x = np.quantile(X, 0.75)
iqr_x = q3_x - q1_x

results1 = pd.DataFrame({
    'Показатель': ['mean', 'std', 'median', 'Q1', 'Q3', 'IQR'],
    'Значение': [mean_x, std_x, median_x, q1_x, q3_x, iqr_x]
})
print("\n", results1.to_string(index=False))

# Графики
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(12, 4))

ax1.hist(X, bins=50, color='steelblue', edgecolor='black', alpha=0.7)
ax1.axvline(mean_x, color='red', linestyle='--', linewidth=2, label=f'mean = {mean_x:.4f}')
ax1.axvline(median_x, color='green', linestyle='--', linewidth=2, label=f'median = {median_x:.4f}')
ax1.set_xlabel('AveBedrms')
ax1.set_ylabel('Частота')
ax1.set_title('Гистограмма X')
ax1.legend()

ax2.boxplot(X, vert=True, patch_artist=True)
ax2.set_ylabel('AveBedrms')
ax2.set_title('Boxplot X')
ax2.grid(axis='y', alpha=0.3)

plt.tight_layout()
plt.show()

# ============================================
# ЗАДАНИЕ 2. РОБАСТНЫЕ ОЦЕНКИ
# ============================================
print("\n" + "=" * 70)
print("ЗАДАНИЕ 2. Робастные оценки")
print("=" * 70)

# Trimmed mean
trim_n = int(len(X) * trim_percent / 100)
sorted_X = np.sort(X)
trimmed_mean_x = np.mean(sorted_X[trim_n:-trim_n])

# Winsorized mean
wins_n = int(len(X) * wins_percent / 100)
winsorized = sorted_X.copy()
winsorized[:wins_n] = sorted_X[wins_n]
winsorized[-wins_n:] = sorted_X[-wins_n-1]
winsorized_mean_x = np.mean(winsorized)

# MAD
mad_x = np.median(np.abs(X - median_x))
scaled_mad_x = 1.4826 * mad_x

# IQR/1.349
iqr_norm_x = iqr_x / 1.349

results2 = pd.DataFrame({
    'Оценка положения': ['mean', 'trimmed mean (10%)', 'winsorized mean (20%)'],
    'Значение': [mean_x, trimmed_mean_x, winsorized_mean_x]
})
print("\n", results2.to_string(index=False))

results2b = pd.DataFrame({
    'Оценка разброса': ['std', 'MAD', '1.4826 * MAD', 'IQR/1.349'],
    'Значение': [std_x, mad_x, scaled_mad_x, iqr_norm_x]
})
print("\n", results2b.to_string(index=False))

# ============================================
# ЗАДАНИЕ 3. ВЛИЯНИЕ ОДНОГО ВЫБРОСА
# ============================================
print("\n" + "=" * 70)
print("ЗАДАНИЕ 3. Влияние одного выброса")
print("=" * 70)

# Создаём X2: max = 15 * median
X2 = X.copy()
max_idx = np.argmax(X2)
old_max = X2[max_idx]
new_max = 15 * median_x
X2[max_idx] = new_max

print(f"\nВыброс: максимальное значение было {old_max:.4f}, стало {new_max:.4f}")
print(f"Медиана = {median_x:.4f}, 15 * медиана = {new_max:.4f}\n")

# Функция для расчёта всех оценок
def get_all_stats(data, trim_n, wins_n):
    sorted_data = np.sort(data)
    return {
        'mean': np.mean(data),
        'std': np.std(data, ddof=1),
        'median': np.median(data),
        'Q1': np.quantile(data, 0.25),
        'Q3': np.quantile(data, 0.75),
        'IQR': np.quantile(data, 0.75) - np.quantile(data, 0.25),
        'trimmed_mean': np.mean(sorted_data[trim_n:-trim_n]),
        'winsorized_mean': np.mean(np.clip(data, sorted_data[trim_n], sorted_data[-trim_n-1])),
        'MAD': np.median(np.abs(data - np.median(data))),
        'scaled_MAD': 1.4826 * np.median(np.abs(data - np.median(data))),
        'IQR_norm': (np.quantile(data, 0.75) - np.quantile(data, 0.25)) / 1.349
    }

stats_X = get_all_stats(X, trim_n, wins_n)
stats_X2 = get_all_stats(X2, trim_n, wins_n)

# Считаем относительные изменения
deltas = {}
for key in stats_X.keys():
    if stats_X[key] != 0:
        deltas[key] = (stats_X2[key] - stats_X[key]) / abs(stats_X[key]) * 100

results3 = pd.DataFrame({
    'Оценка': list(deltas.keys()),
    'Δ (%)': [f"{d:.2f}%" for d in deltas.values()]
})
print(results3.to_string(index=False))

# ============================================
# ЗАДАНИЕ 4. BOOTSTRAP ИНТЕРВАЛЫ 95%
# ============================================
print("\n" + "=" * 70)
print(f"ЗАДАНИЕ 4. Bootstrap 95% доверительные интервалы (B = {B})")
print("=" * 70)

np.random.seed(123)

def bootstrap_ci(data, func, B, alpha=0.05):
    estimates = []
    n = len(data)
    for _ in range(B):
        sample = np.random.choice(data, n, replace=True)
        estimates.append(func(sample))
    lower = np.percentile(estimates, 100 * alpha / 2)
    upper = np.percentile(estimates, 100 * (1 - alpha / 2))
    return lower, upper

# Bootstrap для mean
mean_ci = bootstrap_ci(X, np.mean, B)

# Bootstrap для median
median_ci = bootstrap_ci(X, np.median, B)

# Bootstrap для trimmed mean
trimmed_func = lambda x: np.mean(np.sort(x)[trim_n:-trim_n])
trimmed_ci = bootstrap_ci(X, trimmed_func, B)

results4 = pd.DataFrame({
    'Статистика': ['mean', 'median', f'trimmed mean ({trim_percent}%)'],
    'Нижняя граница (2.5%)': [mean_ci[0], median_ci[0], trimmed_ci[0]],
    'Верхняя граница (97.5%)': [mean_ci[1], median_ci[1], trimmed_ci[1]]
})
print("\n", results4.to_string(index=False))

print("\n" + "=" * 70)
print("ГОТОВО! Все 4 задания выполнены.")
print("=" * 70)