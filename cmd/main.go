package main

import (
	"fmt"
	"livecode/cmd/internal/pkg/storage" // Make sure this is the correct path to your package
	"log"
)

func main() {
	fmt.Println("Логгер")
	storageInstance, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("Ошибка при создании хранилища: %v", err)
	}

	storageInstance.Set("mystring", "myValueString")
	storageInstance.Set("22", "52")
	fmt.Println("=================================")
	fmt.Println("Значения добавлены в хранилище")
	valueS := storageInstance.Get("mystring")
	valueD := storageInstance.Get("22")
	if valueS != nil {
		fmt.Printf("Значение по ключу : %s\n", *valueS)
		fmt.Println(storageInstance.GetKind("mystring")) // Correct: passing key to GetKind
	} else {
		fmt.Println("Значение по ключу не найдено")
	}
	if valueD != nil {
		fmt.Printf("Значение по ключу : %s\n", *valueD)
		fmt.Println(storageInstance.GetKind("22")) // Correct: passing key to GetKind
	} else {
		fmt.Println("Значение по ключу не найдено")
	}

	missingValue := storageInstance.Get("unknownKey")
	if missingValue == nil {
		fmt.Println("Значение по ключу 'unknownKey' не найдено")
	}
}
