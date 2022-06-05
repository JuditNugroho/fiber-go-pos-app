$(function () {
    initTblCarts();
})

function deleteCartFmt() {
    return '<a class="delete-cart" href="javascript:void(0)" title="Hapus"><i class="fa fa-trash"></i></a>'
}

window.eventDeleteCart = {
    'click .delete-cart': function (e, value, row, index) {
        alertify.confirm('Dialog Konfirmasi', 'Apakah anda yakin ingin menghapus data ini?', function () {
            deleteProductToTblCart(index);
        }, null).setting({'labels': {ok: 'Ya', cancel: 'Tidak'}});
    }
}

function initTblCarts() {
    $('#tableCarts').bootstrapTable({
        locale: $('#locale').val(),
        formatNoMatches: function () {
            return '';
        },
        columns: [
            [
                {
                    width: 75,
                    title: 'Del',
                    field: 'action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: deleteCartFmt,
                    events: window.eventDeleteCart,
                },
                {
                    width: 220,
                    align: 'left',
                    field: 'name',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Nama Barang',
                },
                {
                    width: 180,
                    title: 'Harga',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    field: 'price_used',
                    formatter: priceFormatter
                },
                {
                    width: 170,
                    field: 'qty',
                    align: 'center',
                    widthUnit: "px",
                    title: 'Jumlah',
                    valign: 'middle',
                    formatter: qtyFormatter,
                },
                {
                    width: 180,
                    title: 'Total',
                    field: 'total',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    formatter: priceFormatter
                },
            ],
        ]
    });
}

$(document).on('change', "input.input-qty-number", function () {
    // initialize variable
    let qty_value = parseInt($(this).val());
    let minValue = parseInt($(this).attr('min'));
    let maxValue = parseInt($(this).attr('max'));
    let dataIndex = parseInt($(this).attr('data-index'));

    // check if current value is higher than minimum value
    if (qty_value < minValue) {
        $(this).val($(this).data('oldValue'));
        alert('Maaf, qty pembelian melebihi nilai minimum pembelian. Min qty : ' + minValue.toString());
        return false
    }

    // check if current value is smaller than maximum value
    if (qty_value > maxValue) {
        $(this).val($(this).data('oldValue'));
        alert('Maaf, qty pembelian melebihi nilai maksimum pembelian. Max qty : ' + maxValue.toString());
        return false
    }

    // set the new value
    $(this).val($(this).val().replace(/^0+/, ''));
    updateProductToTblCart(dataIndex, qty_value);
});

