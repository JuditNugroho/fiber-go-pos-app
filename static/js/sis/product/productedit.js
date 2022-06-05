function editProduct(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row){
    let ppnFmt = row["ppn"] === true ? "Ya" : "Tidak"
    $("#modalUpsert #product_id").val(row["product_id"]);
    $("#modalUpsert #name").val(row["name"]);
    $("#modalUpsert #barcode").val(row["barcode"]);
    $("#modalUpsert #stock").val(row["stock"]);
    $("#modalUpsert #ppn").val(ppnFmt);
    $("#modalUpsert #price").val(row["price"]);
    $("#modalUpsert #member_price").val(row["member_price"]);
    $("#modalUpsert #discount").val(row["discount"]);
}