package main

import "fmt"

func main() {

    sklad := make (map[string]int)
    for {
        fmt.Println("\nВведите куда вы хотите попасть:\n1. Меню\n2. Выйти")
        var y int
        _, err := fmt.Scan(&y)
        if err != nil {
            fmt.Println("Неверный ввод!")
            return
        }
        if y == 2 {
            fmt.Println("До свидания")
            return}else{
                fmt.Printf("\nВыберите действие:\n1. Добавить товар\n2. Удалить товар\n3. Изменить количество\n" +
                    "4. Проверить наличие\n5. Показать все товары \n6. Выйти\n")
            }
    var x int
    fmt.Scan(&x)
    var name string
    var amount int
    switch {
    case x == 1: 
        fmt.Println("Введите количество товаров, которые хотите добавить")
        var z int
        fmt.Scan(&z)
        fmt.Println("Введите название товаров и их количество через Enter")
        for i:=0; i < z; i++ {
            fmt.Scan(&name ,&amount)
            sklad[name] = amount
        }
        fmt.Println("Товары добавлены!")

    case x == 2:
        fmt.Println("Введите название товара, который нужно удалить")
        fmt.Scan(&name)
        _,exists := sklad[name]
        if exists {
            delete(sklad,name)
            fmt.Printf("Товар %s удален",name)
        }else{
            fmt.Println("Такого товара на складе нет!")
        }
    case x == 3:
        fmt.Println("Введите название товара и новое количество для него")
        fmt.Scan(&name, &amount)
        _, exists := sklad[name]
        if exists {
            sklad[name] = amount
            fmt.Printf("Количество товара обновлено!\n Теперь %s = %d", name,amount)
        }else{
            fmt.Println("Такого товара на складе нет! Возможно Вам нужно выбрать функцию добавить товар.")
        }
    case x == 4:
        fmt.Println("Введите название товара, для которого вы хотите узнать наличие?")
        _,err:=fmt.Scan(&name)
        if err != nil {
            fmt.Println("Неправильный ввод! Попробуйте снова.")
            return
        }
        _, exists := sklad[name]
        if exists {
            fmt.Printf("%s есть на складе в количестве %d",name, sklad[name])
        }else{
            fmt.Printf("%s нет на складе!", name)
        }
    case x == 5:
        fmt.Println("Вот, что есть на складе:")
        count:=0
        for name,amount := range sklad {
            count++
            fmt.Println(count,name, amount)
        }
    case x == 6:
        fmt.Println("Выход")
        return
    default:
        fmt.Println("Неверный выбор! Попробуйте снова!")
    }

    }

}