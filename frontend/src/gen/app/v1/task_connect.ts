// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file app/v1/task.proto (package app.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { TaskServiceCreateRequest, TaskServiceCreateResponse, TaskServiceDeleteRequest, TaskServiceDeleteResponse, TaskServiceGetRequest, TaskServiceGetResponse, TaskServiceListRequest, TaskServiceListResponse, TaskServiceUpdateRequest, TaskServiceUpdateResponse } from "./task_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service app.v1.TaskService
 */
export const TaskService = {
  typeName: "app.v1.TaskService",
  methods: {
    /**
     * @generated from rpc app.v1.TaskService.Create
     */
    create: {
      name: "Create",
      I: TaskServiceCreateRequest,
      O: TaskServiceCreateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc app.v1.TaskService.Get
     */
    get: {
      name: "Get",
      I: TaskServiceGetRequest,
      O: TaskServiceGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc app.v1.TaskService.List
     */
    list: {
      name: "List",
      I: TaskServiceListRequest,
      O: TaskServiceListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc app.v1.TaskService.Update
     */
    update: {
      name: "Update",
      I: TaskServiceUpdateRequest,
      O: TaskServiceUpdateResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc app.v1.TaskService.Delete
     */
    delete: {
      name: "Delete",
      I: TaskServiceDeleteRequest,
      O: TaskServiceDeleteResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

