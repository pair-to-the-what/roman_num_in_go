package hiker

var numToRom = map[int]string {
        0 : "",
        1 : "I",
        2 : "II",
        3 : "III",
        4 : "IV",
        5 : "V",
        6 : "VI",
        9 : "IX" }

func Convert_to_rom(num int) string {
    rom_num := ""
    for num >= 10 {
        rom_num += "X"
        num -= 10
    }
    return rom_num + numToRom[num]
}
