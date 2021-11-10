from app.foobarqix import compute

def test_1_is_1():
    assert compute("1") == "1"

def test_2_is_2():
    assert compute("2") == "2"

def test_3_is_FooFoo():
    assert compute("3") == "FooFoo"

def test_4_is_4():
    assert compute("4") == "4"

def test_5_is_BarBar():
    assert compute("5") == "BarBar"

def test_6_is_Foo():
    assert compute("6") == "Foo"

def test_7_is_QixQix():
    assert compute("7") == "QixQix"

def test_8_is_8():
    assert compute("8") == "8"

def test_9_is_Foo():
    assert compute("9") == "Foo"

def test_10_is_Bar():
    assert compute("10") == "Bar"

def test_13_is_Foo():
    assert compute("13") == "Foo"

def test_15_is_FooBarBar():
    assert compute("15") == "FooBarBar"

def test_21_is_FooQix():
    assert compute("21") == "FooQix"

def test_33_is_FooFooFoo():
    assert compute("33") == "FooFooFoo"

def test_51_is_FooBar():
    assert compute("51") == "FooBar"

def test_53_is_BarFoo():
    assert compute("53") == "BarFoo"
