'use strict';System.register(['../vendor/underscore','../vendor/js-cookie'],function(_export,_context){var isObject,isEmpty,size,Cookie;return {setters:[function(_vendorUnderscore){isObject=_vendorUnderscore.isObject;isEmpty=_vendorUnderscore.isEmpty;size=_vendorUnderscore.size;},function(_vendorJsCookie){Cookie=_vendorJsCookie;}],execute:function(){const memories={};window.addEventListener('storage',_ref => {let key=_ref.key;let newValue=_ref.newValue;if(key in memories){memories[key].cached=parseSet(newValue);}});function parseSet(set){let val;try{val=JSON.parse(set);}catch(e){}return isObject(val)?val:{};}class Memory{constructor(key,expiry){this.key=key;memories[key]=this;this.expiry=expiry;this.cached=this.read();setTimeout(() => this.purgeExpired(),5000);}now(){return Math.floor(Date.now()/1000);}purgeAll(){localStorage.removeItem(this.key);}read(){const key=localStorage.getItem(this.key);if(!key){return {};}return parseSet(key);}has(key){return !!this.chached[key];}writeAll(set){if(isEmpty(set)){return this.purgeAll();}localStorage.setItem(this.key,JSON.stringify(set));}write(key){this.cached=this.read();this.cached[key]=this.now();this.writeAll(this.cached);return size(this.cached);}size(){return size(this.cached);}purgeExpired(){this.chached=this.read(),now=this.now(),limit=86400*this.expiry,expired=[];for(let key in this.chached){if(now>this.chached[key]+limit){expired.push(key);}}if(!expired.length){return;}for(let key of expired){delete this.chached[key];}this.writeAll(this.chached);}}_export('default',Memory);}};});
//# sourceMappingURL=maps/memory.js.map
