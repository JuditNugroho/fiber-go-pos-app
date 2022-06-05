function addProductToTblCart(row) {
    // check if data already inserted
    if (checkDataCartFromProductID(row["product_id"])) {
        alertify.alert('Pesan Dialog', "Barang telah diinputkan ke dalam keranjang");
        return false
    }

    // custom data to support with table
    row["qty"] = 1;
    row["price_used"] = $("#member-id").val() ? row["price"] : row["member_price"];
    row["total"] = row["price_used"] * row["qty"];

    // insert data to some object
    insertDataCart(row);
    $("#is-inserted").val("true");
    $('#tableCarts').bootstrapTable('load', getAllDataCarts());
    setTotalSection()
}

function deleteProductToTblCart(index) {
    $('#tableCarts').bootstrapTable('remove', {
        field: '$index',
        values: [index]
    });
    setTotalSection();
}

function updateProductToTblCart(index, qty_value) {
    updateQtyDataCart(index, qty_value);
    $('#tableCarts').bootstrapTable('updateRow', {
        index: index,
        row: getDataCartByIndex(index),
    });
    setTotalSection();
}

function setTotalSection() {
    // initialization jquery selector
    let totalPPNObj = $("#totalPPN");
    let totalItemObj = $("#totalItem");
    let totalPriceObj = $("#totalPrice");
    let totalDiscountObj = $("#totalDiscount");
    let totalPembelianObj = $("#totalPembelian");

    // initialization temporary obj
    let totalPPN = 0
    let totalItem = 0
    let totalPrice = 0
    let totalDiscount = 0

    $.each(getAllDataCarts(), function (index, value) {
        totalItem += value["qty"]
        totalPrice += value["total"]
        totalDiscount += value["discount"] * value["qty"]

        if (value.ppn) {
            totalPPN += (value["total"] * ppnDecimal)
        }
    });

    totalPPNObj.html(formatIDR(totalPPN));
    totalItemObj.html(formatQty(totalItem));
    totalPriceObj.html(formatIDR(totalPrice));
    totalDiscountObj.html(formatIDR(totalDiscount));

    let totalPembelian = totalPrice + totalPPN - totalDiscount;
    totalPembelianObj.html(formatIDR(totalPembelian));
}

function resetCurrentTransaction() {
    removeAllDataCart()
    $('#tableCarts').bootstrapTable('load', getAllDataCarts());
    // initialization jquery selector

    $("#totalPPN").html("0");
    $("#totalItem").html("0");
    $("#totalPrice").html("0");
    $("#totalDiscount").html("0");
    $("#totalPembelian").html("0");
}