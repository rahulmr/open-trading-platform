/* eslint-disable */
// source: static-data-service.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var listing_pb = require('./listing_pb.js');
goog.object.extend(proto, listing_pb);
goog.exportSymbol('proto.staticdataservice.ListingId', null, global);
goog.exportSymbol('proto.staticdataservice.ListingIds', null, global);
goog.exportSymbol('proto.staticdataservice.Listings', null, global);
goog.exportSymbol('proto.staticdataservice.MatchParameters', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.staticdataservice.ListingId = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.staticdataservice.ListingId, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.staticdataservice.ListingId.displayName = 'proto.staticdataservice.ListingId';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.staticdataservice.ListingIds = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.staticdataservice.ListingIds.repeatedFields_, null);
};
goog.inherits(proto.staticdataservice.ListingIds, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.staticdataservice.ListingIds.displayName = 'proto.staticdataservice.ListingIds';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.staticdataservice.Listings = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.staticdataservice.Listings.repeatedFields_, null);
};
goog.inherits(proto.staticdataservice.Listings, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.staticdataservice.Listings.displayName = 'proto.staticdataservice.Listings';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.staticdataservice.MatchParameters = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.staticdataservice.MatchParameters, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.staticdataservice.MatchParameters.displayName = 'proto.staticdataservice.MatchParameters';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.staticdataservice.ListingId.prototype.toObject = function(opt_includeInstance) {
  return proto.staticdataservice.ListingId.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.staticdataservice.ListingId} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.ListingId.toObject = function(includeInstance, msg) {
  var f, obj = {
    listingid: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.staticdataservice.ListingId}
 */
proto.staticdataservice.ListingId.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.staticdataservice.ListingId;
  return proto.staticdataservice.ListingId.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.staticdataservice.ListingId} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.staticdataservice.ListingId}
 */
proto.staticdataservice.ListingId.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setListingid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.staticdataservice.ListingId.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.staticdataservice.ListingId.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.staticdataservice.ListingId} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.ListingId.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getListingid();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
};


/**
 * optional int32 listingId = 1;
 * @return {number}
 */
proto.staticdataservice.ListingId.prototype.getListingid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.staticdataservice.ListingId.prototype.setListingid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.staticdataservice.ListingIds.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.staticdataservice.ListingIds.prototype.toObject = function(opt_includeInstance) {
  return proto.staticdataservice.ListingIds.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.staticdataservice.ListingIds} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.ListingIds.toObject = function(includeInstance, msg) {
  var f, obj = {
    listingidsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.staticdataservice.ListingIds}
 */
proto.staticdataservice.ListingIds.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.staticdataservice.ListingIds;
  return proto.staticdataservice.ListingIds.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.staticdataservice.ListingIds} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.staticdataservice.ListingIds}
 */
proto.staticdataservice.ListingIds.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Array<number>} */ (reader.readPackedInt32());
      msg.setListingidsList(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.staticdataservice.ListingIds.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.staticdataservice.ListingIds.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.staticdataservice.ListingIds} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.ListingIds.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getListingidsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
};


/**
 * repeated int32 listingIds = 1;
 * @return {!Array<number>}
 */
proto.staticdataservice.ListingIds.prototype.getListingidsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/** @param {!Array<number>} value */
proto.staticdataservice.ListingIds.prototype.setListingidsList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 */
proto.staticdataservice.ListingIds.prototype.addListingids = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.staticdataservice.ListingIds.prototype.clearListingidsList = function() {
  this.setListingidsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.staticdataservice.Listings.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.staticdataservice.Listings.prototype.toObject = function(opt_includeInstance) {
  return proto.staticdataservice.Listings.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.staticdataservice.Listings} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.Listings.toObject = function(includeInstance, msg) {
  var f, obj = {
    listingsList: jspb.Message.toObjectList(msg.getListingsList(),
    listing_pb.Listing.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.staticdataservice.Listings}
 */
proto.staticdataservice.Listings.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.staticdataservice.Listings;
  return proto.staticdataservice.Listings.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.staticdataservice.Listings} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.staticdataservice.Listings}
 */
proto.staticdataservice.Listings.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new listing_pb.Listing;
      reader.readMessage(value,listing_pb.Listing.deserializeBinaryFromReader);
      msg.addListings(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.staticdataservice.Listings.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.staticdataservice.Listings.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.staticdataservice.Listings} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.Listings.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getListingsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      listing_pb.Listing.serializeBinaryToWriter
    );
  }
};


/**
 * repeated model.Listing listings = 1;
 * @return {!Array<!proto.model.Listing>}
 */
proto.staticdataservice.Listings.prototype.getListingsList = function() {
  return /** @type{!Array<!proto.model.Listing>} */ (
    jspb.Message.getRepeatedWrapperField(this, listing_pb.Listing, 1));
};


/** @param {!Array<!proto.model.Listing>} value */
proto.staticdataservice.Listings.prototype.setListingsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.model.Listing=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.Listing}
 */
proto.staticdataservice.Listings.prototype.addListings = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.model.Listing, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.staticdataservice.Listings.prototype.clearListingsList = function() {
  this.setListingsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.staticdataservice.MatchParameters.prototype.toObject = function(opt_includeInstance) {
  return proto.staticdataservice.MatchParameters.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.staticdataservice.MatchParameters} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.MatchParameters.toObject = function(includeInstance, msg) {
  var f, obj = {
    symbolmatch: jspb.Message.getFieldWithDefault(msg, 1, ""),
    namematch: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.staticdataservice.MatchParameters}
 */
proto.staticdataservice.MatchParameters.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.staticdataservice.MatchParameters;
  return proto.staticdataservice.MatchParameters.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.staticdataservice.MatchParameters} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.staticdataservice.MatchParameters}
 */
proto.staticdataservice.MatchParameters.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSymbolmatch(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNamematch(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.staticdataservice.MatchParameters.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.staticdataservice.MatchParameters.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.staticdataservice.MatchParameters} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.staticdataservice.MatchParameters.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSymbolmatch();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getNamematch();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string symbolMatch = 1;
 * @return {string}
 */
proto.staticdataservice.MatchParameters.prototype.getSymbolmatch = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.staticdataservice.MatchParameters.prototype.setSymbolmatch = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string nameMatch = 2;
 * @return {string}
 */
proto.staticdataservice.MatchParameters.prototype.getNamematch = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.staticdataservice.MatchParameters.prototype.setNamematch = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.staticdataservice);
