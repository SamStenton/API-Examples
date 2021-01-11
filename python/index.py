import requests, ast
# Use given API url to perform a get request and return text of response
# Prints status code and returns false if not 200 status code
def getRequest():
	x = requests.get('https://engineering-task.elancoapps.com/api/applications')
	if (x.status_code != 200):
		print("status code" + x.status_code)
		return False
	else:
		return x.text
# Decodes a given request from a string into an Array
def decodeRequest(req):
	arr = ast.literal_eval(req)
	return arr

# Takes in an array and prints all elements of the array
def printArray(arr):
	for var in arr:
		print (var)

#
def main():
	response = getRequest()
	if (response):
		printArray(decodeRequest(response))
	else:
		print("there was an error")
main()