require_relative "../common/mock_vars"
require_relative "../common/command_helpers"
require_relative "../common/mock_helper"
require_relative "../common/trace_response_builder"
require_relative "spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "dp "
  return cb
end

RSpec.describe 'Workspace test cases', :type => :aruba, :feature => "Workspaces", :severity => :critical do
  include_context "shared command helpers"
  include_context "mock shared vars"

  before(:each) do
    MockHelper.resetMock()
  end

  it "Workspace - List", :story => "List Workspaces", :severity => :normal, :testId => 1 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/workspaces/workspaces.json")

      expectedEndpointResponse = TraceResponseBuilder.getWorkspacesResponseFactory(responseHash)
      MockHelper.setupResponse("getWorkspaces", responseHash)

      result = cb.workspace.list.build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Workspace - Describe", :story => "Describe Workspaces", :severity => :normal, :testId => 2 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/workspaces/workspace.json")

      expectedEndpointResponse = TraceResponseBuilder.getWorkspaceByNameResponseFactory(responseHash)
      MockHelper.setupResponse("getWorkspaceByName", responseHash)

      result = cb.workspace.describe.name("mock@hortonworks.com").build(false)
      resultHash = MockHelper.getResultHash(result.output)

      #expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end
end
