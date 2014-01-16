SameTimeTrackStatus
===================

Logs the status of users in IBM Lotus Notes 8+


Setup:

1. Lotus Notes must be installed.
2. startWebContainer must be set to true to enable the webapi.

    File: `C:\Program Files\lotus\notes\framework\rcp\plugin_customization.ini`
    
    Set `com.ibm.collaboration.realtime.webapi/startWebContainer=true`

3. Lotus Notes must be running.

Run:

    go run main.go
    
Output:

	Logging userId123
	1: Status at 2014-01-16 12:35:33.609375 -0600 CST
	2: Status at 2014-01-16 12:40:33.609375 -0600 CST
	3: Status at 2014-01-16 12:45:33.609375 -0600 CST
	
Also there will be a file called `status.txt` that will store all the responses from the webapi.

WebAPI documentation:

- [IBM Sametime Software Developer Kit (SDK)](http://www14.software.ibm.com/webapp/download/nochargesearch.jsp?q0=&k=ALL&S_TACT=104CBW71&status=Active&b=Lotus&sr=1&q=sametime+sdk&ibm-search=Search)
