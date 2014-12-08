package hiker

import ("testing")

var romTests = []struct {
        arabic int
        roman  string
}{
        {1, "I"},
        {2, "II"},
        {3, "III"},
        {4, "IV"},
        {10, "X"},
        {11, "XI"},
        {15, "XV"},
        {29, "XXIX"},
        {36, "XXXVI"},
        {40, "XL"},
}

func assertConversion(t *testing.T, num int, expected string) {
    got := Convert_to_rom(num)
    if got != expected {
        t.Error("convert failed for", num, "expected", expected, "got", got)
    }
}
    
func Test_all(t *testing.T) {
    for _, pair := range romTests {
        assertConversion(t, pair.arabic, pair.roman)
    }
}
