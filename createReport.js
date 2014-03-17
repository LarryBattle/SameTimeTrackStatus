var x = [
// Array of timestamp objects.
]
var filterById = function (id) {
	return x.filter(function (o) {
		return o.username === id;
	});
};
var getTimeAndStatus = function (arr) {
	return arr.map(function (o) {
		return "time: " + o.timestamp + ", " + o.status;
	});
};
var getTimeAndStatusDiff = function (arr) {
	var lastStatus = -1;
	return arr.filter(function (o) {
		if (lastStatus == o.status) {
			return false;
		}
		lastStatus = o.status;
		return true;
	}).map(function (o) {
		return "time: " + o.timestamp + ", " + (o.statusMessage || "Left Office");
	});
};
var getProps = function (arr, prop) {
	return arr.map(function (o) {
		return o[prop];
	});
};
var getWeeklyStatusOn = function (id) {
	var userTS = filterById(id);
	console.log(userTS[0].displayName);
	console.log(getTimeAndStatusDiff(userTS).join("\n"));
};

var getUniqueIds = function () {
	var ids = {};
	for (var i = 0, l = x.length; i < l; i++) {
		if (!ids[x[i].username]) {
			ids[x[i].username] = 1;
		}
	}
	return Object.keys(ids);
};
clear();
getUniqueIds().forEach(getWeeklyStatusOn);
