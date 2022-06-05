$(function () {
    initTable();
})


function actionFormatter() {
    return [
        '<a class="edit" href="javascript:void(0)" title="Edit"><i class="fa fa-edit"></i></a>',
        '<a class="remove" href="javascript:void(0)" title="Remove"><i class="fa fa-trash"></i></a>'
    ].join('')
}

window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editProduct(row);
    },
    'click .remove': function (e, value, row, index) {
        alertify.confirm('Dialog Konfirmasi', 'Apakah anda yakin ingin menghapus data ini?',
            function () {
                deleteProduct(row);
            }, null).setting({'labels': {ok: 'Ya', cancel: 'Tidak'}});
    }
}

function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'ID Barang',
                    field: 'product_id',
                },
                {
                    width: 350,
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Nama Barang',
                },
                {
                    width: 250,
                    align: 'left',
                    widthUnit: "px",
                    title: 'Barcode',
                    field: 'barcode',
                    valign: 'middle',
                },
                {
                    width: 150,
                    title: 'Stok',
                    field: 'stock',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 100,
                    title: 'PPN',
                    field: 'ppn',
                    widthUnit: "px",
                    align: 'center',
                    valign: 'middle',
                    formatter: checkboxFormatter
                },
                {
                    width: 200,
                    title: 'Harga',
                    field: 'price',
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    formatter: priceFormatter
                },
                {
                    width: 200,
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Harga Member',
                    field: 'member_price',
                    formatter: priceFormatter
                },
                {
                    width: 200,
                    align: 'center',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Diskon',
                    field: 'discount',
                    formatter: priceFormatter
                },
                {
                    width: 200,
                    title: 'Action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: actionFormatter,
                    events: window.eventActions,
                }
            ],
        ]
    });
}

// your custom ajax request here
function ajaxRequest(params) {
    let page = 1;
    let req = params.data;
    let baseURL = $('#baseURL').text();
    if (params.data["offset"] !== 0) {
        page = (params.data["offset"] / params.data["limit"]) + 1
    }

    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/dt_products",
        'contentType': 'application/json',
        "data": {
            "page": page,
            "limit": req["limit"],
            "search": req["search"],
        },
    }).done(function (data) {
        $('#table').bootstrapTable('resetView');
        params.success(data);
    });
}

function clearFormInput() {
    $("#modalUpsert #product_id").val("");
    $("#modalUpsert #name").val("");
    $("#modalUpsert #barcode").val("0");
    $("#modalUpsert #stock").val("0");
    $("#modalUpsert #ppn").val("Ya");
    $("#modalUpsert #price").val("0");
    $("#modalUpsert #member_price").val("0");
    $("#modalUpsert #discount").val("0");
}