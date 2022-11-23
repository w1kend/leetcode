class Solution:
    max = 214748364

    def reverse(self, x: int) -> int:
        if x < 0:
            sign = -1
            max_last_digit = 8
        else:
            sign = 1
            max_last_digit = 7

        reversed_str = str(x)[::-1].removesuffix('-')

        # if the number is definitely fits into 32 bytes
        if len(reversed_str) < 9:
            return sign * int(reversed_str)

        last_digit = int(reversed_str[len(reversed_str)-1:])
        reversed_save = int(reversed_str[:len(reversed_str)-1])

        if reversed_save < self.max or reversed_save == self.max and last_digit <= max_last_digit:
            return sign * int(reversed_str)

        return 0
