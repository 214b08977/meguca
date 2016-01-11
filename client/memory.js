import {isObject, isEmpty, size} from '../vendor/underscore'
import * as Cookie from '../vendor/js-cookie'

// All instances of the Memory class
const memories = {}

/**
 * Listen for storage events and update the stored value for exising memory
 * instances, if the key changes. These only fire, if the write happens in
 * another tab of the same origin.
 */
window.addEventListener('storage', ({key, newValue}) => {
	if (key in memories) {
		memories[key].cached = parseSet(newValue)
	}
})

/**
 * Parse a stringified set
 * @param {string} set
 * @returns {Object}
 */
function parseSet(set) {
	let val
	try {
		val = JSON.parse(set)
	}
	catch(e) {}
	return isObject(val) ? val : {}
}

/**
 * Self-expiring localStorage set manager
 */
export default class Memory {
	/**
	 * Construct a new localStorage controller
	 * @param {string} key - localStorage key
	 * @param {int} expiry - Entry lifetime in days
	 */
	constructor(key, expiry) {
		this.key = key
		memories[key] = this
		this.expiry = expiry

		// Read the initial value
		this.cached = this.read()

		// Purge old entries on start
		setTimeout(() => this.purgeExpired(), 5000)
	}

	/**
	 * Return current time in seconds
	 * @returns {int}
	 */
	now() {
		return Math.floor(Date.now() / 1000)
	}

	/**
	 * Clear the stored set
	 */
	purgeAll() {
		localStorage.removeItem(this.key)
	}

	/**
	 * Read and parse the stringified set from localStorage
	 * @returns {Object}
	 */
	read() {
		const key = localStorage.getItem(this.key)
		if (!key) {
			return {}
		}
		return parseSet(key)
	}

	/**
	 * Return, if the given jey exists in the set
	 * @param {string} key
	 * @returns {bool}
	 */
	has(key) {
	    return !!this.chached[key]
	}

	/**
	 * Replace the existing set, if any, with the suplied one
	 * @param {Object} object
	 */
	writeAll(set) {
		if (isEmpty(set)) {
			return this.purgeAll()
		}
		localStorage.setItem(this.key, JSON.stringify(set))
	}

	/**
	 * Write a single key to the stored set
	 * @returns {int} - Size of new set
	 */
	write(key) {
		// When performing writes, best fetch everything, rather than rely on
		// events for browser tab cache synchronisation. Browser backround tab
		// optimisation might fuck us over.
		this.cached = this.read()
		this.cached[key] = this.now()
		this.writeAll(this.cached)
		return size(this.cached) // Return number of items
	}

	/**
	 * Return the current size of the stored Set
	 */
	size() {
		return size(this.cached)
	}

	/**
	 * Delete expired entries from set and write to localStorage
	 */
	purgeExpired() {
		this.chached = this.read(),
			now = this.now(),
			limit = 86400 * this.expiry,
			expired = []
		for (let key in this.chached) {
			if (now > this.chached[key] + limit) {
				expired.push(key)
			}
		}
		if (!expired.length) {
			return
		}
		for (let key of expired) {
		    delete this.chached[key]
		}
		this.writeAll(this.chached)
	}
}
