class Solution:
    values = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    substructions = {
        'I': {
            'V': True,
            'X': True,
        },
        'X': {
            'L': True,
            'C': True,
        },
        'C': {
            'D': True,
            'M': True,
        },
    }

    def romanToInt(self, s: str) -> int:
        sum = 0
        l = len(s)-1
        i = 0
        while i <= l:
            letter = s[i]
            val = self.values[letter]
            if val == 0:
                return 0

            if i < l and i+1 <= l:
                # check next value
                nletter = s[i+1]
                nval = self.values[nletter]
                if nval > val and letter in self.substructions:
                    substr = self.substructions[letter]
                    if nletter in substr:
                        sum += nval - val
                        i += 2
                        continue

            sum += val
            i += 1

        return sum
