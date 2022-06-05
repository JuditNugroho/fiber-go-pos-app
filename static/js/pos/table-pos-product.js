$(function () {
    initTblProduct();
})

function addToCartFmt() {
    return '<a class="add-cart" href="javascript:void(0)" title="Tambah Barang"><i class="fa fa-shopping-cart"></i></a>'
}

window.addCartEventAction = {
    'click .add-cart': function (e, value, row, index) {
        addProductToTblCart(row)
    }
}

function initTblProduct() {
    $('#tableProducts').bootstrapTable({
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
                    width: 150,
                    title: 'Action',
                    field: 'action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: addToCartFmt,
                    events: window.addCartEventAction,
                }
            ],
        ]
    });
}

function headerTblProduct(column) {
    return {
        product_id: {
            css: {
                background: '#34b9f9',
                color : '#FFFFFF'
            },
        },
        name: {
            css: {
                background: '#34b9f9',
                color : '#FFFFFF'
            }
        },
        action: {
            css: {
                background: '#34b9f9',
                color : '#FFFFFF'
            }
        }
    }[column.field]
}

// your custom ajax request here
function ajaxProductRequest(params) {
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
