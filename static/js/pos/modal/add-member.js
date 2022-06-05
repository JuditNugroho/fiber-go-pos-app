function addMember() {
    if ($("#is-inserted").val() === "true") {
        alertify.alert('Pesan Dialog', "Data pembelian sudah diinputkan, Tidak bisa memasukkan member");
        return false
    }
    if ($('#btnAddMember').prop('disabled')){
        return false
    }
    $('#modalAddMember').modal('show');
    $('#table-add-member').bootstrapTable('resetView')
}