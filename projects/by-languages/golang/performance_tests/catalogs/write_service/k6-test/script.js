/*
 * Catalogs Write-Service Api
 * Catalogs Write-Service Api.
 *
 * OpenAPI spec version: 1.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://github.com/OpenAPITools/openapi-generator
 *
 * OpenAPI generator version: 6.1.0-SNAPSHOT
 */


import http from "k6/http";
import { group, check, sleep } from "k6";

const BASE_URL = "";
// Sleep duration between successive requests.
// You might want to edit the value of this variable or remove calls to the sleep function on the script.
const SLEEP_DURATION = 0.1;
// Global variables should be initialized.

export default function() {
    group("/api/v1/products/search", () => {
        let search = 'TODO_EDIT_THE_SEARCH'; // specify value as there is no example value for this parameter in OpenAPI spec
        let size = 'TODO_EDIT_THE_SIZE'; // specify value as there is no example value for this parameter in OpenAPI spec
        let orderBy = 'TODO_EDIT_THE_ORDERBY'; // specify value as there is no example value for this parameter in OpenAPI spec
        let page = 'TODO_EDIT_THE_PAGE'; // specify value as there is no example value for this parameter in OpenAPI spec

        // Request No. 1
        {
            let url = BASE_URL + `/api/v1/products/search?orderBy=${orderBy}&page=${page}&search=${search}&size=${size}`;
            let request = http.get(url);

            check(request, {
                "OK": (r) => r.status === 200
            });
        }
    });

    group("/api/v1/products", () => {
        let size = 'TODO_EDIT_THE_SIZE'; // specify value as there is no example value for this parameter in OpenAPI spec
        let orderBy = 'TODO_EDIT_THE_ORDERBY'; // specify value as there is no example value for this parameter in OpenAPI spec
        let page = 'TODO_EDIT_THE_PAGE'; // specify value as there is no example value for this parameter in OpenAPI spec

        // Request No. 1
        {
            let url = BASE_URL + `/api/v1/products?orderBy=${orderBy}&page=${page}&size=${size}`;
            let request = http.get(url);

            check(request, {
                "OK": (r) => r.status === 200
            });

            sleep(SLEEP_DURATION);
        }

        // Request No. 2
        {
            let url = BASE_URL + `/api/v1/products`;
            // TODO: edit the parameters of the request body.
            let body = {"description": "string", "name": "string", "price": "bigdecimal"};
            let params = {headers: {"Content-Type": "application/json", "Accept": "application/json"}};
            let request = http.post(url, JSON.stringify(body), params);

            check(request, {
                "Created": (r) => r.status === 201
            });
        }
    });

    group("/api/v1/products/{id}", () => {
        let id = 'TODO_EDIT_THE_ID'; // specify value as there is no example value for this parameter in OpenAPI spec

        // Request No. 1
        {
            let url = BASE_URL + `/api/v1/products/${id}`;
            let request = http.get(url);

            check(request, {
                "OK": (r) => r.status === 200
            });

            sleep(SLEEP_DURATION);
        }

        // Request No. 2
        {
            let url = BASE_URL + `/api/v1/products/${id}`;
            let request = http.del(url);

            check(request, {
                "": (r) => r.status === 204
            });
        }
    });

}
