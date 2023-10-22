class Stack():

    def __init__(self, values=None):
        self.buf = []
        if values:
            for val in values:
                self.put(val)

    def put(self, value):
        self.buf.append(value)

    def pop(self):
        return self.buf.pop()

    def peek(self):
        if self.buf:
            return self.buf[-1]
        return None

    def __len__(self):
        return len(self.buf)

    def __bool__(self):
        return bool(self.buf)

    def __str__(self):
        return f'{self.buf}'
