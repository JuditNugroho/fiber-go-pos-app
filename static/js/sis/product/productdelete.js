function deleteProduct(row) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        type: "POST",
        async: false,
        data: JSON.stringify(row),
        contentType: 'application/json',
        url: baseURL + "svc/product/delete",
    }).then(function (res) {
        $('#table').bootstrapTable('refresh');
        alertify.success("Data barang berhasil dihapus");
    }).catch(function (a) {
        alertify.error("Error : " + a.responseText);
    });
}