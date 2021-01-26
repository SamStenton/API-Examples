$Url = 'https://engineering-task.elancoapps.com/api/applications' 
#API Call
$result = Invoke-RestMethod -Method Get -Uri $Url 
#Printing out the results
foreach($result in $results){
    Write-Output($result)
}