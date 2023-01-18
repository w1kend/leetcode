class BrowserHistory:
    history: list
    i: int
    len: int

    def __init__(self, homepage: str):
        self.history = [homepage]
        self.i = 0
        self.len = 1

    def visit(self, url: str) -> None:
        if self.i == len(self.history) - 1:
            self.history.append(url)
            self.len += 1
            self.i += 1
        else:
            self.i += 1
            self.history[self.i] = url
            self.len = self.i+1

    def back(self, steps: int) -> str:
        idx = self.i - steps
        if idx < 0:
            idx = 0
        self.i = idx
        return self.history[idx]

    def forward(self, steps: int) -> str:
        idx = self.i + steps
        if idx >= self.len-1:
            idx = self.len-1
        self.i = idx
        return self.history[idx]
