/** https://github.com/talentmaker/site/blob/master/src/utils/request.ts */

import {inlineTryPromise} from "@luke-zhang-04/utils"
import {phraseStatuses} from "@luke-zhang-04/utils/http"

type ThenArg<T> = T extends PromiseLike<infer U> ? U : T
type Methods = "GET" | "HEAD" | "POST" | "PUT" | "DELETE" | "CONNECT" | "OPTIONS" | "TRACE"
type Conversions = "arrayBuffer" | "blob" | "formData" | "json" | "text"

type ConversionTypes = {[P in Exclude<Conversions, "json">]: ThenArg<ReturnType<Body[P]>>} & {
    json: {[key: string]: unknown}
}

class HTTPError extends Error {
    public readonly name = "HTTPError"

    public constructor(
        public readonly status: number,
        message: string,
    ) {
        super(`${phraseStatuses[status as keyof typeof phraseStatuses] ?? status} - ${message}`)
    }
}

type Config<T> = {
    /** Method of request e.g `GET`, `POST` */
    method: Methods

    /** What to convert the response to e.g `json`, `text` */
    conversion?: T

    /** Request body */
    body?: {[key: string]: unknown}

    /** Request headers */
    authorization?: string

    /** Content type e.g `application/json`, `multipart/form-data` */
    contentType?: string

    /** Other request params e.g `{headers: {credentials: "include"}}` */
    init?: RequestInit
}

/**
 * Wrapper for the fetch API that is hopefully a bit more intuitive
 *
 * @param url - Url of request
 * @param method - Method of request e.g `GET`, `POST`
 * @param conversion - What to convert the response to e.g `json`, `text`
 * @param body - Request body
 * @param authorization - Access token for Authorization header
 * @param contentType - Content type e.g `application/json`, `multipart/form-data`
 * @param init - Other request params e.g `{headers: {credentials: "include"}}`
 * @returns The converted response
 */
export async function request<T extends Conversions>(
    url: string,
    method: Methods,
    conversion: T,
    body?: {[key: string]: unknown},
    authorization?: string,
    contentType?: string,
    init?: RequestInit,
): Promise<ConversionTypes[T]>

/**
 * Wrapper for the fetch API that is hopefully a bit more intuitive
 *
 * @param url - Url of request
 * @param method - Method of request e.g `GET`, `POST`
 * @param conversion - What to convert the response to e.g `json`, `text`
 * @param body - Request body
 * @param authorization - Access token for Authorization header
 * @param contentType - Content type e.g `application/json`, `multipart/form-data`
 * @param init - Other request params e.g `{headers: {credentials: "include"}}`
 * @returns A raw response
 */
export async function request(
    url: string,
    method: Methods,
    conversion?: undefined,
    body?: {[key: string]: unknown},
    authorization?: string,
    contentType?: string,
    init?: RequestInit,
): Promise<Response>

/**
 * Wrapper for the fetch API that is hopefully a bit more intuitive
 *
 * @param url - Url of request
 * @param config - Request config
 * @returns The converted response
 */
export async function request<T extends Conversions>(
    url: string,
    config: Config<T>,
): Promise<ConversionTypes[T]>

/**
 * Wrapper for the fetch API that is hopefully a bit more intuitive
 *
 * @param url - Url of request
 * @param config - Request config
 * @returns A raw response
 */
export async function request(url: string, config: Config<undefined>): Promise<Response>

/** Wrapper for the fetch API that is hopefully a bit more intuitive */
export async function request<T extends Conversions>(
    url: string,
    methodOrConfig: Methods | Config<T>,
    conversion?: T,
    body?: {[key: string]: unknown},
    authorization?: string,
    contentType?: string,
    init: RequestInit = {},
): Promise<Response | ConversionTypes[T]> {
    const config =
        typeof methodOrConfig === "string"
            ? {
                  method: methodOrConfig,
                  conversion,
                  body,
                  contentType,
                  authorization,
                  headers: init.headers,
                  init,
              }
            : {
                  ...methodOrConfig,
                  headers: init.headers,
              }

    const response = await fetch(url, {
        ...config.init,
        method: config.method,
        headers: {
            ...config.headers,
            "Content-Type": config.contentType ?? "application/json",
            ...(config.authorization ? {Authorization: `Bearer ${config.authorization}`} : {}),
        },
        body: config.body ? JSON.stringify(config.body) : undefined,
    })

    if (!response.ok) {
        const msg = await inlineTryPromise<{[key: string]: unknown}>(() => response.json(), false)

        if (msg) {
            throw new HTTPError(
                response.status,
                typeof msg.message === "string" ? msg.message : JSON.stringify(msg),
            )
        }

        throw new HTTPError(response.status, response.statusText)
    }

    if (config.conversion) {
        return await response[config.conversion]()
    }

    return response
}

export default request
