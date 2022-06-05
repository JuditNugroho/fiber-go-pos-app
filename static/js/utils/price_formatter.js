function priceFormatter(value) {
    return formatIDR(value)
}

function formatIDR(number) {
    if (number === 0) {
        return '0'
    }

    if (number <= 0) {
        return 'Rp. -' + formatPrice(Math.abs(number).toString())
    }

    let number_string = number.toString(),
        split = number_string.split(','),
        sisa = split[0].length % 3,
        rupiah = split[0].substr(0, sisa),
        ribuan = split[0].substr(sisa).match(/\d{3}/gi);

    // tambahkan titik jika yang di input sudah menjadi angka ribuan
    let separator;
    if (ribuan) {
        separator = sisa ? '.' : '';
        rupiah += separator + ribuan.join('.');
    }

    rupiah = split[1] !== undefined ? rupiah + ',' + split[1] : rupiah;
    return 'Rp. ' + rupiah;
}

function formatPrice(number) {
    if (number === 0) {
        return '0'
    }

    let number_string = number.replace(/[^,\d]/g, "").toString(),
        split = number_string.split(','),
        sisa = split[0].length % 3,
        rupiah = split[0].substr(0, sisa),
        ribuan = split[0].substr(sisa).match(/\d{3}/gi);

    // tambahkan titik jika yang di input sudah menjadi angka ribuan
    let separator;
    if (ribuan) {
        separator = sisa ? '.' : '';
        rupiah += separator + ribuan.join('.');
    }

    rupiah = split[1] !== undefined ? rupiah + ',' + split[1] : rupiah;
    return rupiah;
}

function convertIDRToNumber(rupiah) {
    return parseInt(rupiah.replace(/[^,\d]/g, ""));
}