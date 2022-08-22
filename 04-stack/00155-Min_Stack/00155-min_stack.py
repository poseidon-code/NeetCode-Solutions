# 155: Min stack
# https://leetcode.com/problems/min-stack/


# SOLUTION
class MinStack:
    def __init__(self):
        self.s1 = []
        self.s2 = []
    
    def push(self, x: int)  -> None:
        self.s1.append(x)
        if len(self.s2)==0 or x <= self.getMin():
            self.s2.append(x)
    
    def pop(self) -> None:
        if self.s1[-1] == self.getMin():
            self.s2.pop()
        self.s1.pop()
    
    def top(self) -> int:
        return self.s1[-1]
    
    def getMin(self) -> int:
        return self.s2[-1]



if __name__ == "__main__":
    o = MinStack()
    
    # OPERATIONS
    o.push(-2)
    o.push(0)
    o.push(-3)
    print(o.getMin(), end=" ")
    o.pop()
    o.top()
    print(o.getMin(), end=" ")

    print()