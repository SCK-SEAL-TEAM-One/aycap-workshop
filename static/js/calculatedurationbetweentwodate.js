$(function () {
    $('#calculate').click(GetApi)
})

function GetApi() {
    var host = "http://localhost:3000/CalculateDurationBetweenTwoDate"
    var parameter = `?start_day=${$('#start_day').val()}` +
        `&start_month=${$('#start_month').val()}` +
        `&start_year=${$('#start_year').val()}` +
        `&end_day=${$('#end_day').val()}` + 
        `&end_month=${$('#end_month').val()}` + 
        `&end_year=${$('#end_year').val()}`
    var url = host + parameter
    
    $.getJSON(url, function (responseData) {
        $('#from').html(responseData['from'])
        $('#to').html(responseData['to'])
        $('#days').html(responseData['days'])
        $('#seconds').html(responseData['seconds'])
        $('#minutes').html(responseData['minutes'])
        $('#hours').html(responseData['hours'])
    })
}