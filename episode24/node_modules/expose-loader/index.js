/*
	MIT License http://www.opensource.org/licenses/mit-license.php
	Author Tobias Koppers @sokra
*/

var path = require('path');

function accesorString(value) {
	var childProperties = value.split(".");
	var length = childProperties.length;
	var propertyString = "global";
	var result = "";

	for(var i = 0; i < length; i++) {
		if(i > 0)
			result += "if(!" + propertyString + ") " + propertyString + " = {};\n";
		propertyString += "[" + JSON.stringify(childProperties[i]) + "]";
	}

	result += "module.exports = " + propertyString;
	return result;
}

module.exports = function() {};
module.exports.pitch = function(remainingRequest) {
	// Change the request from an /abolute/path.js to a relative ./path.js
	// This prevents [chunkhash] values from changing when running webpack
	// builds in different directories.
	const newRequestPath = remainingRequest.replace(
		this.resourcePath,
		'.' + path.sep + path.relative(this.context, this.resourcePath)
	);
	this.cacheable && this.cacheable();
	if(!this.query) throw new Error("query parameter is missing");
    /*
     * Workaround until module.libIdent() in webpack/webpack handles this correctly.
     *
     * fixes:
     * - https://github.com/webpack-contrib/expose-loader/issues/55
     * - https://github.com/webpack-contrib/expose-loader/issues/49
     */
	this._module.userRequest = this._module.userRequest + '-exposed';
	return accesorString(this.query.substr(1)) + " = " +
		"require(" + JSON.stringify("-!" + newRequestPath) + ");";
};
