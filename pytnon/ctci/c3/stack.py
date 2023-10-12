class Stack():

    def __init__(self):
        self.buf = []

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
