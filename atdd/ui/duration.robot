***Setting***
Library    SeleniumLibrary
Test Teardown     Close All Browsers

***Variables***
${URL}    http://localhost:3000/web
${Browser}    chrome

***Test Cases***
เข้าเว็บไซต์กรอกข้อมูลวันที่เริ่ม 7/3/1977 วันที่สิ้นสุด 21/6/2018
    Open Browser    ${URL}    ${Browser}
    Title Should Be    Calculate Duration Between Two Dates
    ##กรอกวันที่เริ่มต้น
    Input Text    id=start_day    7
    ##กรอกเดือนเริ่มต้น
    Input Text    id=start_month   3
    ##กรอกปีเริ่มต้น
    Input Text    id=start_year    1977 
    ##กรอกวันที่สิ้นสุด
    Input Text    id=end_day    21
    ##กรอกเดือนสิ้นสุด
    Input Text    id=end_month   6
    ##กรอกปีสิ้นสุด
    Input Text    id=end_year    2018 
    ##กดปุ่มคำนวณหาผลลัพธ์
    Click Element    id=calculate
    ##แสดงจำนวนผลลัพธ์ระหว่างวันที่เริ่มต้นและวันที่สิ้นสุด
    Wait Until Page Contains    15,082 days