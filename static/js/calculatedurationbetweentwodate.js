$(function () {
    $('#calculate').click(GetApi)
})

function GetApi() {
    var host = "http://localhost:3000/CalculateDurationBetweenTwoDate"
    var parameter = `?startday=${$('#startDay').val()}` +
        `&startmonth=${$('#startMonth').val()}` +
        `&startyear=${$('#startYear').val()}` +
        `&endday=${$('#endDay').val()}` + 
        `&endmonth=${$('#endMonth').val()}` + 
        `&endyear=${$('#endYear').val()}`
    var url = host + parameter
    
    $.getJSON(url, function (responseData) {
        $('#FromAndIncluding').html(responseData['From and including'])
        $('#ToAndIncluding').html(responseData['To and including'])
        $('#Result').html(responseData['Result'])
        $('#Seconds').html(responseData['Seconds'])
        $('#Minutes').html(responseData['Minutes'])
        $('#Hours').html(responseData['Hours'])
        $('#WeeksAndDays').html(responseData['WeeksAndDays'])
    })
}