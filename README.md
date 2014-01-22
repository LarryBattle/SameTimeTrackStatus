SameTimeTrackStatus
===================

Logs the status of users in IBM Lotus Notes SameTime 8+


##Setup:

1. [Google Go](http://golang.org/) must be installed.
2. Lotus Notes must be installed.
3. startWebContainer must be set to true to enable the webapi.

    File: `%ProfileFiles%\lotus\notes\framework\rcp\plugin_customization.ini`
    
    Set `com.ibm.collaboration.realtime.webapi/startWebContainer=true`

4. Lotus Notes must be running.

##Run:

    go run main.go -userid=USER_ID

##Dev Documentation:

- [IBM Sametime Software Developer Kit (SDK)](http://www14.software.ibm.com/webapp/download/nochargesearch.jsp?q0=&k=ALL&S_TACT=104CBW71&status=Active&b=Lotus&sr=1&q=sametime+sdk&ibm-search=Search)

##Note:

- This has only been tested in Windows XP.
- Users can disable `Auto-Status Changes` in Lotus Notes for SameTime. Thus preventing others from knowing when a user is away from their desk.
