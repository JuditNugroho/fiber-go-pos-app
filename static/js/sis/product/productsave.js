$("#upsertProduct").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        alertify.alert('Pesan Peringatan', param.err);
        return
    }

    saveProduct(param.data);
});

function getParamValue() {
    let product_id = $("#modalUpsert #product_id").val().trim()
    let name = $("#modalUpsert #name").val().trim()
    let barcode = $("#modalUpsert #barcode").val().trim()
    let stock = parseInt($("#modalUpsert #stock").val().trim())
    let ppn = $("#modalUpsert #ppn").val().trim().trim() === "Ya"
    let price = parseFloat($("#modalUpsert #price").val().trim())
    let member_price = parseFloat($("#modalUpsert #member_price").val().trim())
    let discount = parseFloat($("#modalUpsert #discount").val().trim())

    if (product_id === "") {
        return {"data": null, "err": "ID barang tidak boleh kosong !"}
    }

    if (name === "") {
        return {"data": null, "err": "Nama barang tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "product_id": product_id,
            "name": name,
            "barcode": barcode,
            "stock": stock,
            "ppn": ppn,
            "price": price,
            "member_price": member_price,
            "discount": discount
        }
    }

}

function saveProduct(data) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        type: "POST",
        async: false,
        data: JSON.stringify(data),
        contentType: 'application/json',
        url: baseURL + "svc/product/upsert",
    }).then(function (res) {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (a) {
        alertify.alert('Pesan Error', a.responseText);
    });
}