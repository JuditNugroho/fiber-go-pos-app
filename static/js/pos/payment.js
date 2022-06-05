function openPaymentModal() {
    if (getTotalDataCarts() === 0) {
        alertify.alert('Pesan Dialog', "Harap lakukan transaksi terlebih dahulu");
        return false
    }
    $('#modalPayment').modal('show');
    $('#modalPayment #totalTransaction').val($("#totalPembelian").text());
    $("#modalPayment #totalPayment").val("0");
    $("#modalPayment #totalPayment").trigger("change")

    let totalPayment = convertIDRToNumber($("#modalPayment #totalPayment").val());
    let totalTransaction = convertIDRToNumber($("#modalPayment #totalTransaction").val());
    let change = totalPayment - totalTransaction;
    $("#modalPayment #numeric_change").val(change);
    $("#modalPayment #change").val(formatIDR(change));
}

function buttonConsumerMoney(money) {
    $("#modalPayment #totalPayment").val(money)
    $("#modalPayment #totalPayment").trigger("change")
}

$("#modalPayment #totalPayment").on("change", function () {
    $(this).val(formatPrice($(this).val()))

    let totalPayment = convertIDRToNumber($(this).val());
    let totalTransaction = convertIDRToNumber($("#modalPayment #totalTransaction").val());

    let change = totalPayment - totalTransaction
    $("#modalPayment #numeric_change").val(change)
    $("#modalPayment #change").val(formatIDR(change))
});

$("#modalPayment #totalPayment").on("keydown", function (e) {
    // Allow: backspace, delete, tab, escape, enter and .
    if ($.inArray(e.keyCode, [46, 8, 9, 27, 13]) !== -1 ||
        // Allow: Ctrl+A
        (e.keyCode === 65 && e.ctrlKey === true) ||
        // Allow: home, end, left, right
        (e.keyCode >= 35 && e.keyCode <= 39)) {
        // let it happen, don't do anything
        return;
    }
    // Ensure that it is a number and stop the keypress
    if ((e.shiftKey || (e.keyCode < 48 || e.keyCode > 57)) && (e.keyCode < 96 || e.keyCode > 105)) {
        e.preventDefault();
    }
});

function processPayment() {
    let change = parseFloat($("#modalPayment #numeric_change").val().trim());
    if (change < 0) {
        alert("Transaksi invalid. Uang yang anda masukkan kurang dari total pembelian");
        return false
    }
    alertify.success("Transaksi Berhasil");
    $('#modalPayment').modal('hide');
    resetCurrentTransaction()
}