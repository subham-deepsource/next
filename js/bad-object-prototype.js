const test_object = {
    'company': 'DeepSource',
    'conference': 'DeepSource Next\'21',
    'date': 'June 8, 2021',
}

const hasConferenceName = test_object.hasOwnProperty('conference')

if(hasConferenceName){
    console.log('Welcome to', test_object['conference'])
}