import chartmogul
config = chartmogul.Config('api key')

# notes = chartmogul.Customer.notes(config, uuid='cus_208bb082-9f05-11ee-972a-23cb260bb1e2')
# print(notes)
# note = chartmogul.Customer.createNote(config, uuid='cus_208bb082-9f05-11ee-972a-23cb260bb1e2', data={
#     'text': 'This is a note',
#     'type': 'note'
# })
# print(note)
# new_note = chartmogul.CustomerNote.create(config, data={
#     'customer_uuid': 'cus_208bb082-9f05-11ee-972a-23cb260bb1e2',
#     'text': 'uhhh',
#     'type': 'note'
# })
# print(new_note)
# notes = chartmogul.CustomerNote.all(config, per_page=20, customer_uuid='cus_208bb082-9f05-11ee-972a-23cb260bb1e2')
# print(notes)
# target_note = chartmogul.CustomerNote.retrieve(config, uuid='note_bfd9f816-9f4a-11ee-8365-a3fb386a58f5')
# print(target_note)
# updated_note = chartmogul.CustomerNote.patch(config, uuid='note_bfd9f816-9f4a-11ee-8365-a3fb386a58f5', data={
#     'text': 'This is a new note'
# })
# print(updated_note)
deleted_note = chartmogul.CustomerNote.destroy(config, uuid='note_bfd9f816-9f4a-11ee-8365-a3fb386a58f5')
print(deleted_note)
