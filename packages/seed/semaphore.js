/**
 * A JS implementation of a semaphore is easy because everything in JS is thread-safe (JS is single
 * threaded, Promises are fake concurrency).
 *
 * @param promises - Array of functions that return a promise
 * @param limit - Max number of concurrent promises that may run
 * @returns Array of results, with no guarantees for ordering
 */
export const semaphore = (promises, limit) =>
    new Promise((resolve, reject) => {
        const results = []
        const running = new Array(limit)
        let inputIndex = 0 // Index from the input array
        let runningCount = 0
        let didReject = false

        const thenFunc = (index) => (result) => {
            results.push(result)

            if (inputIndex >= promises.length) {
                runningCount--
                delete running[index]

                if (runningCount === 0) {
                    resolve(results)
                }
            } else if (!didReject) {
                running[index] = promises[inputIndex]()
                    .then(thenFunc(inputIndex))
                    .catch((err) => {
                        didReject = true
                        reject(err)
                    })

                inputIndex++
            }

            return result
        }

        for (; inputIndex < limit && inputIndex < promises.length - 1; inputIndex++) {
            runningCount++
            running[inputIndex] = promises[inputIndex]()
                .then(thenFunc(inputIndex))
                .catch((err) => {
                    didReject = true
                    reject(err)
                })
        }
    })
