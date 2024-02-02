import { AccountsAPI, Configuration, axiosRetry, Paginator, SearchAPI, TransformsAPI, TransformsApiCreateTransformRequest, Search, IdentityDocument} from "sailpoint-api-client"

const createTransform = async () => {

    let apiConfig = new Configuration()
    let api = new TransformsAPI(apiConfig)
    let transform: TransformsApiCreateTransformRequest = 
    {
        transform:
        {
            name: "Test Transform",
            type: "dateFormat",
            attributes: {
                inputFormat: "MMM dd yyyy, HH:mm:ss.SSS",
                outputFormat: "yyyy/dd/MM"
            }
        }
    }
    const val = await api.createTransform(transform)
    console.log(val)
}

const search = async () => {
    let apiConfig = new Configuration()
    let api = new SearchAPI(apiConfig)
    let search: Search = {
        indices: [
            "identities"
        ],
        query: {
            query: "*"
        },
        sort: ["-name"]
	}
    const val = await Paginator.paginateSearchAPI(api, search, 100, 1000)

    for (const result of val.data) {
        const castedResult: IdentityDocument = result
        console.log(castedResult.name)
    }
    
}

const getPaginatedAccounts = async () => {

    
    let apiConfig = new Configuration()
    apiConfig.retriesConfig = {
        retries: 4,
        retryDelay: axiosRetry.exponentialDelay,
        onRetry(retryCount, error, requestConfig) {
            console.log(`retrying due to request error, try number ${retryCount}`)
        },
    }
    let api = new AccountsAPI(apiConfig)
    
    const val = await Paginator.paginate(api, api.listAccounts, {limit: 100}, 10)

    console.log(val)

}


const getPaginatedTransforms = async () => {

    
    let apiConfig = new Configuration()
    apiConfig.retriesConfig = {
        retries: 4,
        retryDelay: axiosRetry.exponentialDelay,
        onRetry(retryCount, error, requestConfig) {
            console.log(`retrying due to request error, try number ${retryCount}`)
        },
    }
    let api = new TransformsAPI(apiConfig)
    
    const val = await Paginator.paginate(api, api.listTransforms, {limit: 250}, 100)

    console.log(val.data.length)

}



getPaginatedTransforms()