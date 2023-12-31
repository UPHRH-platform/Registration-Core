// Code generated by go-swagger; DO NOT EDIT.

package upload_and_create_records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// PostV1SchemaNameUploadOKCode is the HTTP code returned for type PostV1SchemaNameUploadOK
const PostV1SchemaNameUploadOKCode int = 200

/*
PostV1SchemaNameUploadOK OK

swagger:response postV1SchemaNameUploadOK
*/
type PostV1SchemaNameUploadOK struct {

	/*
	  In: Body
	*/
	Payload models.CreateRecordResponse `json:"body,omitempty"`
}

// NewPostV1SchemaNameUploadOK creates PostV1SchemaNameUploadOK with default headers values
func NewPostV1SchemaNameUploadOK() *PostV1SchemaNameUploadOK {

	return &PostV1SchemaNameUploadOK{}
}

// WithPayload adds the payload to the post v1 schema name upload o k response
func (o *PostV1SchemaNameUploadOK) WithPayload(payload models.CreateRecordResponse) *PostV1SchemaNameUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 schema name upload o k response
func (o *PostV1SchemaNameUploadOK) SetPayload(payload models.CreateRecordResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1SchemaNameUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostV1SchemaNameUploadBadRequestCode is the HTTP code returned for type PostV1SchemaNameUploadBadRequest
const PostV1SchemaNameUploadBadRequestCode int = 400

/*
PostV1SchemaNameUploadBadRequest Bad Request

swagger:response postV1SchemaNameUploadBadRequest
*/
type PostV1SchemaNameUploadBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewPostV1SchemaNameUploadBadRequest creates PostV1SchemaNameUploadBadRequest with default headers values
func NewPostV1SchemaNameUploadBadRequest() *PostV1SchemaNameUploadBadRequest {

	return &PostV1SchemaNameUploadBadRequest{}
}

// WithPayload adds the payload to the post v1 schema name upload bad request response
func (o *PostV1SchemaNameUploadBadRequest) WithPayload(payload *models.ErrorPayload) *PostV1SchemaNameUploadBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 schema name upload bad request response
func (o *PostV1SchemaNameUploadBadRequest) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1SchemaNameUploadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1SchemaNameUploadNotFoundCode is the HTTP code returned for type PostV1SchemaNameUploadNotFound
const PostV1SchemaNameUploadNotFoundCode int = 404

/*
PostV1SchemaNameUploadNotFound Not found

swagger:response postV1SchemaNameUploadNotFound
*/
type PostV1SchemaNameUploadNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewPostV1SchemaNameUploadNotFound creates PostV1SchemaNameUploadNotFound with default headers values
func NewPostV1SchemaNameUploadNotFound() *PostV1SchemaNameUploadNotFound {

	return &PostV1SchemaNameUploadNotFound{}
}

// WithPayload adds the payload to the post v1 schema name upload not found response
func (o *PostV1SchemaNameUploadNotFound) WithPayload(payload *models.ErrorPayload) *PostV1SchemaNameUploadNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 schema name upload not found response
func (o *PostV1SchemaNameUploadNotFound) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1SchemaNameUploadNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1SchemaNameUploadInternalServerErrorCode is the HTTP code returned for type PostV1SchemaNameUploadInternalServerError
const PostV1SchemaNameUploadInternalServerErrorCode int = 500

/*
PostV1SchemaNameUploadInternalServerError Internal Server Error

swagger:response postV1SchemaNameUploadInternalServerError
*/
type PostV1SchemaNameUploadInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewPostV1SchemaNameUploadInternalServerError creates PostV1SchemaNameUploadInternalServerError with default headers values
func NewPostV1SchemaNameUploadInternalServerError() *PostV1SchemaNameUploadInternalServerError {

	return &PostV1SchemaNameUploadInternalServerError{}
}

// WithPayload adds the payload to the post v1 schema name upload internal server error response
func (o *PostV1SchemaNameUploadInternalServerError) WithPayload(payload *models.ErrorPayload) *PostV1SchemaNameUploadInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 schema name upload internal server error response
func (o *PostV1SchemaNameUploadInternalServerError) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1SchemaNameUploadInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
