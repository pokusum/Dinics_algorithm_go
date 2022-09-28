# Dinic's algorithm
## О алгоритме
[Алгоритм Диница](https://en.wikipedia.org/wiki/Dinic%27s_algorithm) — [полиномиальный алгоритм](https://en.wikipedia.org/wiki/Time_complexity#Strongly_and_weakly_polynomial_time) для нахождения [максимального потока](https://en.wikipedia.org/wiki/Maximum_flow_problem) в [транспортной сети](https://en.wikipedia.org/wiki/Flow_network), предложенный в 1970 году математиком Ефимом Диницем. Временная сложность алгоритма составляет $O(V^2E)$.

## Input and output
### Input:
В первой строке ожидается 4 числа через пробел: $V$ - количество вершин, $E$ - количество рёбер, $start$ - индекс вершины истока, $finish$ - индекс вершины стока.

Дальше ожидается $E$ строчек по 3 числа: $u$ - начало ребра, $v$ - конец ребра, $cap$ - пропускная способность ребра.

### Output:
Программа выведет одно число $maxFlow$ являющееся мощностью максимального потока.