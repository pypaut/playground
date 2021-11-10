def compute(input_str):
    result = ""
    input_int = int(input_str)

    if input_int % 3 == 0:
        result += "Foo"
    if input_int % 5 == 0:
        result += "Bar"
    if input_int % 7 == 0:
        result += "Qix"
    elif len(input_str) == 1 and input_int % 3 and input_int % 5 and input_int % 7:
        result += input_str

    for c in input_str:
        c_int = int(c)
        if c_int == 3:
            result += "Foo"
        elif c_int == 5:
            result += "Bar"
        elif c_int == 7:
            result += "Qix"

    return result
