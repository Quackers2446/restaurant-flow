/**
 * Delay the execution of a function, and reset said delay if the function is called again within
 * the delay window.
 *
 * @example
 *
 * ```ts
 * window.addEventListener(
 *     "resize",
 *     debounce((event) => console.log(event), 100),
 * )
 * ```
 *
 * @param func - Function to debounce
 * @param ms - Time to wait
 * @returns Debounced function
 */
export const debounce = <Args extends unknown[]>(
    func: (...args: Args) => void,
    ms: number,
): ((...args: Args) => void) => {
    let timeout: ReturnType<typeof setTimeout> | undefined = undefined

    return (...args: Args) => {
        clearTimeout(timeout)

        timeout = setTimeout(() => {
            func(...args)
        }, ms)
    }
}
