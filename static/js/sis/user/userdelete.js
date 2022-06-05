function deleteUser(row) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        type: "POST",
        async: false,
        data: JSON.stringify(row),
        contentType: 'application/json',
        url: baseURL + "svc/user/delete",
    }).then(function (res) {
        $('#table').bootstrapTable('refresh');
        alertify.success("Data user berhasil dihapus");
    }).catch(function (a) {
        alertify.error("Error : " + a.responseText);
    });
}