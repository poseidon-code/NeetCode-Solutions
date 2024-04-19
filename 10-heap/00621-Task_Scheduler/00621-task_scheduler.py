# 621: Task Scheduler
# https://leetcode.com/problems/task-scheduler


class Solution:
    # SOLUTION
    def leastInterval(self, tasks: list[str], n: int) -> int:
        counts = [0] * 26
        for task in tasks:
            counts[ord(task) - ord('A')] += 1
        counts.sort()

        chunk = counts[25] - 1
        time = chunk * n

        for i in range(24, -1, -1):
            time -= min(chunk, counts[i])

        return len(tasks) + time if time >= 0 else len(tasks)


if __name__ == "__main__":
    o = Solution()

    # INPUT
    tasks = ['A','C','A','B','D','B']
    n = 1

    # OUTPUT
    result = o.leastInterval(tasks, n)
    print(result)
