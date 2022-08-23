# 150: Evaluate Reverse Polish Notation
# https://leetcode.com/problems/evaluate-reverse-polish-notation/


class Solution:
    # SOLUTION
    def evalRPN(self, tokens: list[str]) -> int:
        result = []

        for s in tokens:
            if len(s)>1 or s[0].isdigit():
                result.append(int(s))
            else:
                x2 = result.pop();
                x1 = result.pop();
                
                if s[0]=='+': x1+=x2
                if s[0]=='-': x1-=x2
                if s[0]=='*': x1*=x2
                if s[0]=='/': x1/=x2

                result.append(x1)

        return result[-1]



if __name__ == "__main__":
    o = Solution()

    # INPUT
    tokens: list[str] = ["2","1","+","3","*"]

    # OUTPUT
    result = o.evalRPN(tokens)
    print(result)