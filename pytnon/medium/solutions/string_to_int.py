class Solution:
    min = -2147483648
    max = 2147483647

    def myAtoi(self, s: str) -> int:
        # trim spaces
        s = s.lstrip()

        if len(s) == 0:
            return 0

        # check for the sign
        sign = 1
        if s[0] == '-':
            sign = -1

        if s[0] == '-' or s[0] == '+':
            s = s[1:]

        # return if the string starts with no digit
        if len(s) > 0 and not s[0].isdigit():
            return 0

        # calculate the number
        number = 0
        for c in s:
            if c.isdigit():
                number = number * 10 + int(c)
            else:
                break

        # check bounds
        if number < self.min or number > self.max:
            return self.min if sign < 0 else self.max

        return sign * number
