<?php
require('./vendor/autoload.php');

ChartMogul\Configuration::getDefaultConfiguration()
    ->setApiKey();

print_r(ChartMogul\Ping::ping()->data);

$customer = ChartMogul\Customer::retrieve('cus_208bb082-9f05-11ee-972a-23cb260bb1e2');
// print customer info
// print_r($customer->data);
print_r($customer_notes = $customer->notes(['page' => 1]));

// $customer = ChartMogul\Customer::retrieve($uuid);
// $new_customer_note = $customer->createNote([
//   "type" => "call",
// ], [
//   'text' => 'This is a note'
// ]);
// print_r($new_customer_note->data);

// $customer_notes = ChartMogul\CustomerNote::all([
//     'customer_uuid' => 'cus_208bb082-9f05-11ee-972a-23cb260bb1e2',
// ])

// $customer_note = ChartMogul\CustomerNote::create([
//     'customer_uuid' => 'cus_208bb082-9f05-11ee-972a-23cb260bb1e2',
//     'type' => 'note',
//     'text' => 'This is a note'
// ])
// $note_uuid = $customer_note.data.uuid
// ChartMogul\CustomerNote::retrieve("note_fd017c92-9f71-11ee-b59b-eb527fbda734");
// ChartMogul\CustomerNote::retrieve("note_fd017c92-9f71-11ee-b59b-eb527fbda734");

// $customer_note = ChartMogul\CustomerNote::create([
//     'customer_uuid' => 'cus_208bb082-9f05-11ee-972a-23cb260bb1e2',
//     'type' => 'note',
//     'text' => 'This is a note'
// ])

// $updated_customer_note = ChartMogul\CustomerNote::update([
//     'note_uuid' => 'note_fd017c92-9f71-11ee-b59b-eb527fbda734',
// ], [
//     'text' => 'This is a new note'
//   ]);

?>
