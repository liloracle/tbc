/**
 * Is this a lookup object generated by Typescript, for a Typescript enum
 * generated by protobuf-ts?
 *
 * - No `const enum` (enum must not be inlined, we need reverse mapping).
 * - No string enum (we need int32 for protobuf).
 * - Must have a value for 0 (otherwise, we would need to support custom default values).
 */
export function isEnumObject(arg) {
    if (typeof arg != 'object' || arg === null) {
        return false;
    }
    if (!arg.hasOwnProperty(0)) {
        return false;
    }
    for (let k of Object.keys(arg)) {
        let num = parseInt(k);
        if (!Number.isNaN(num)) {
            // is there a name for the number?
            let nam = arg[num];
            if (nam === undefined)
                return false;
            // does the name resolve back to the number?
            if (arg[nam] !== num)
                return false;
        }
        else {
            // is there a number for the name?
            let num = arg[k];
            if (num === undefined)
                return false;
            // is it a string enum?
            if (typeof num !== 'number')
                return false;
            // do we know the number?
            if (arg[num] === undefined)
                return false;
        }
    }
    return true;
}
/**
 * Lists all values of a Typescript enum, as an array of objects with a "name"
 * property and a "number" property.
 *
 * Note that it is possible that a number appears more than once, because it is
 * possible to have aliases in an enum.
 *
 * Throws if the enum does not adhere to the rules of enums generated by
 * protobuf-ts. See `isEnumObject()`.
 */
export function listEnumValues(enumObject) {
    if (!isEnumObject(enumObject))
        throw new Error("not a typescript enum object");
    let values = [];
    for (let [name, number] of Object.entries(enumObject))
        if (typeof number == "number")
            values.push({ name, number });
    return values;
}
/**
 * Lists the names of a Typescript enum.
 *
 * Throws if the enum does not adhere to the rules of enums generated by
 * protobuf-ts. See `isEnumObject()`.
 */
export function listEnumNames(enumObject) {
    return listEnumValues(enumObject).map(val => val.name);
}
/**
 * Lists the numbers of a Typescript enum.
 *
 * Throws if the enum does not adhere to the rules of enums generated by
 * protobuf-ts. See `isEnumObject()`.
 */
export function listEnumNumbers(enumObject) {
    return listEnumValues(enumObject)
        .map(val => val.number)
        .filter((num, index, arr) => arr.indexOf(num) == index);
}
