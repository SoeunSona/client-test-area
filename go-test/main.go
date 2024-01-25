package main

import (
	"fmt"
	"os"

	cm "github.com/chartmogul/chartmogul-go/v4"
)

func main() {

	os.Setenv("KEY", "")
	api := cm.API{
		ApiKey: os.Getenv("KEY"),
	}

	// Try authentication
	ok, err := api.Ping()
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}
	fmt.Println(ok, "ok")

	customer_uuid := "cus_208bb082-9f05-11ee-972a-23cb260bb1e2"

	// params := &cm.ListNotesParams{}

	// note_list, err := api.ListCustomerNotes(params, customer_uuid)
	// if err != nil {
	// 	fmt.Printf("This didn't work out: %v", err)
	// }
	// fmt.Printf("%+v\n", note_list)
	note, err := api.CreateCustomerNote(&cm.NewNote{
		Text: "This is a note",
		Type: "note",
	}, customer_uuid)
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}
	fmt.Printf("%+v\n", note)
	new_note, err := api.CreateNote(&cm.NewNote{
		Text:         "This is a note",
		Type:         "note",
		CustomerUUID: customer_uuid,
	})
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}
	fmt.Printf("%+v\n", new_note)

	target_note, err := api.RetrieveNote(new_note.UUID)
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}

	fmt.Printf("Target note: %v", target_note)

	updated_note, err := api.UpdateNote(&cm.UpdateNote{
		Text: "This is an updated note",
	}, new_note.UUID)
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}
	fmt.Println(updated_note)

	list_note, err := api.ListNotes(&cm.ListNotesParams{
		CustomerUUID: customer_uuid,
	})
	if err != nil {
		fmt.Printf("This didn't work out: %v", err)
	}
	fmt.Println(list_note)

	api.DeleteNote(new_note.UUID)
	// if error != nil {
	// 	fmt.Printf("This didn't work out: %v", error)
	// }
	// str := strconv.FormatBool(hi)
	// fmt.Println(str)
}
