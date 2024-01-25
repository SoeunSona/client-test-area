const ChartMogul = require('chartmogul-node');
const config = new ChartMogul.Config();

async function main() {
    const customerUuid = 'cus_208bb082-9f05-11ee-972a-23cb260bb1e2';
    const notes = await ChartMogul.Customer.notes(config, customerUuid, { per_page: 10 })
    console.log(notes)
    const new_note = await ChartMogul.Customer.createNote(config, customerUuid, {
        text: 'This is a note',
        type: 'note'
    })
    console.log(new_note)

   const new_note1 = await ChartMogul.CustomerNote.create(config, {
        customer_uuid: customerUuid,
        type: 'note',
        text: 'This is a note'
    })
    console.log(new_note1)

    const anote = await ChartMogul.CustomerNote.retrieve(config, new_note1.uuid)
    console.log(anote)

    const update_note = await ChartMogul.CustomerNote.patch(config, new_note1.uuid, {
        text: 'This is an updated note'
    })
    console.log(update_note)

    const destroy_note = await ChartMogul.CustomerNote.destroy(config, new_note1.uuid)
    console.log(destroy_note)
}

main().then(() => {
    console.log('done')
}).catch((err) => {
    console.log(err)
})

// const new_note1 = ChartMogul.CustomerNote.create(config, {
//     customer_uuid: customerUuid,
//     type: 'note',
//     text: 'This is a note'
// })
// console.log(new_note1)
// const anote = ChartMogul.CustomerNote.retrieve(config)
// console.log(anote)
// const updateChartMogul.CustomerNote.patch(config, noteUuid)
// ChartMogul.CustomerNote.destroy(config, noteUuid)
// ChartMogul.CustomerNote.all(config, { per_page: 10, cursor: 'cursor==', customer_uuid: customerUuid})


