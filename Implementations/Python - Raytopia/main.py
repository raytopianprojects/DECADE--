from dataclasses import dataclass


# Open file that will be interpreted
def file(target_file):
    try:
        with open(fr"{target_file}.dcj", "r") as opened_file:
            target_file = opened_file.read()
        return target_file
    except:
        print(f"ERROR target FILE {target_file} not found! Check entered file name")


@dataclass
class stack_obj:
    stack = {}

    def push(self, var):
        self.stack.update(var)

    def set_value(self, value):
        self.stack.update(value)


def interpreter(file_name):
    # Initialize the stack and the parser
    parse = file(file_name)
    parse = parse.splitlines()
    stacks = stack_obj()

    for x in range(0, len(parse)):
        parse[x] = parse[x].split(" ")

        if parse[x][0][1:] == "DECADE":
            stacks.push({f"{parse[x][1][:-1]}": 0})

        elif parse[x][0][1] == "!":
            print(f"{stacks.stack[parse[x][1][:-1]]}")

        elif parse[x][0][1:] in stacks.stack:
            parse[x][1:] = [" ".join(parse[x][1:])]
            parse[x][1] = parse[x][1].replace('"', '')
            parse[x][1] = parse[x][1].replace("]", '')

            stacks.set_value({f"{parse[x][0][1:]}": parse[x][1:][-1]})

    # print(stacks.stack)


if __name__ == "__main__":
    interpreter("hello_world")
    input()
