/*! (C) 2017 Andrea Giammarchi - MIT Style License */
var Class = (function (Object) {'use strict';
  var
    reserved = ['constructor', 'extends', 'static'],
    isNative = function (fn) {
      return /\[native code\]/.test(Object.toString.call(fn));
    },
    ObjectProto = Object.prototype,
    gPO = Object.getPrototypeOf ||
          function (o) { return o.__proto__; },
    sPO = Object.setPrototypeOf ||
          function (o, p) { o.__proto__ = p; return o; },
    gOPD = Object.getOwnPropertyDescriptor,
    defineProperty = Object.defineProperty,
    defProps = Object.defineProperties,
    hasReflect = typeof Reflect === 'object',
    construct = hasReflect &&
          Reflect.construct ||
          function (Super, args, Constructor) {
            [].unshift.call(args, Super);
            var C = Super.bind.apply(Super, args);
            return sPO(new C, Constructor.prototype);
          },
    ownKeys = hasReflect &&
          Reflect.ownKeys ||
          (Object.getOwnPropertySymbols ?
            function ownKeys(o) {
              return Object.getOwnPropertySymbols(o)
                     .concat(Object.getOwnPropertyNames(o));
            } :
            Object.getOwnPropertyNames),
    superPropertyDescriptor = {
      get: function () {
        var
          Super = function () {
            var
              constructor = self.constructor,
              parent = gPO(proto),
              result
            ;
            sPO(self, parent);
            while ((parent.constructor === constructor)) parent = gPO(parent);
            self.constructor = constructor;
            try {
              result = isNative(parent.constructor) ?
                construct(parent.constructor, arguments, constructor) :
                parent.constructor.apply(self, arguments);
            } finally {
              sPO(self, proto);
            }
            delete self.constructor;
            return result ? sPO(result, proto) : self;
          },
          self = this,
          proto = gPO(self),
          parent = proto,
          call = Super.call,
          apply = Super.apply
        ;
        if (proto) {
          sPO(Super, proto);
          Super.call = call;
          Super.apply = apply;
          do {
            ownKeys(parent).forEach(function (prop) {
              if (Super.hasOwnProperty(prop)) return;
              var
                descriptor = gOPD(parent, prop),
                method = descriptor.value
              ;
              if (typeof method === 'function') {
                descriptor.value = function () {
                  var result, parent = proto;
                  do { parent = gPO(parent); }
                  while ((method === parent[prop]));
                  try { result = parent[prop].apply(sPO(self, parent), arguments); }
                  finally { sPO(self, proto); }
                  return result;
                };
                defineProperty(Super, prop, descriptor);
              }
            });
          } while ((parent = gPO(parent)) && parent !== ObjectProto);
        }
        return Super;
      }
    },
    superStaticDescriptor = {
      get: function () {
        var Super = {}, self = this, parent = self;
        do {
          ownKeys(parent).forEach(function (prop) {
            var
              descriptor = gOPD(parent, prop),
              method = descriptor.value
            ;
            if (typeof method === 'function') {
              descriptor.value = function () {
                var
                  proto = gPO(self), method = self[prop],
                  result, parent = proto
                ;
                while ((method === parent[prop])) parent = gPO(parent);
                self.method = parent[prop];
                try { result = self.method.apply(sPO(self, gPO(parent)), arguments); }
                finally { sPO(self, proto).method = method; }
                return result;
              };
              defineProperty(Super, prop, descriptor);
            }
          });
        } while((parent = gPO(parent)) && parent !== Function);
        return Super;
      }
    }
  ;
  return function Classtrophobic(definition) {
    var
      Constructor = definition.constructor,
      Statics = definition['static'],
      Super = definition['extends'],
      Class = definition.hasOwnProperty('constructor') ?
        function () {
          return Constructor.apply(this, arguments) || this;
        } :
        (Super ?
          (isNative(Super) ?
            function () {
              return construct(Super, arguments, Class);
            }   :
            function () {
              var result = Super.apply(this, arguments);
              return result ? sPO(result, Class.prototype) : this;
            }
          ) :
          function () {}
        ),
      Static = {'super': superStaticDescriptor},
      Prototype = {'super': superPropertyDescriptor}
    ;
    if (Super) {
      sPO(Class, Super);
      sPO(Class.prototype, Super.prototype);
    }
    ownKeys(definition)
      .forEach(function (key) {
        if (reserved.indexOf(key) < 0) {
          Prototype[key] = gOPD(definition, key);
          Prototype[key].enumerable = false;
        }
      });
    defProps(Class.prototype, Prototype);
    if (Statics) ownKeys(Statics).forEach(function (key) {
      Static[key] = gOPD(Statics, key);
      Static[key].enumerable = false;
    });
    Class.prototype.constructor = Class;
    return defProps(Class, Static);
  };
})(Object);
try { module.exports = Class; } catch(o_O) {}
