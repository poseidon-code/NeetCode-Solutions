// 167: Two Sum II
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

class Solution {
    // SOLUTION
    public int[] twoSum(int[] numbers, int target) {
        int l = 0;
        int r = numbers.length - 1;
        int[] result = new int[2];
        
        while (l<r) {
            int s = numbers[l] + numbers[r];
            if (s==target) {
                result[0] = l+1;
                result[1] = r+1;
                break;
            }
            else if (s<target) l++;
            else r--;
        }

        return result;
    }


    public static void main(String[] args) {
        Solution o = new Solution();

        // INPUT
        int[] numbers = {2,7,11,15};
        int target = 9;

        // OUTPUT
        var result = o.twoSum(numbers, target);
        System.out.print("["); for (var v : result) System.out.print(v+" "); System.out.println("\b]");
    }
}